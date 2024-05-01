package tq

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	run "runtime"
	"strings"

	"encoding/json"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/client"
)

// import "github.com/skysyzygy/tq/cmd"

// Global config options and structures for tq
type TqConfig struct {
	// the Tessitura API client
	*client.TessituraServiceWeb

	// Basic auth for requests
	basicAuth runtime.ClientAuthInfoWriter

	// TODO: Bearer token for requests
	// tokenAuth func(*runtime.ClientOperation)

	// Logger, exported so that logging can happen from within the
	// command scripts
	Log *slog.Logger

	// some flags, set by New
	verbose bool
	dryRun  bool
}

func New(logFile *os.File, verbose bool, dryRun bool) *TqConfig {
	logLevel := new(slog.LevelVar)
	if verbose {
		logLevel.Set(slog.LevelInfo)
	} else {
		logLevel.Set(slog.LevelWarn)
	}
	return &TqConfig{
		Log:     slog.New(NewLogHandler(logFile, logLevel)),
		verbose: verbose,
		dryRun:  dryRun,
	}

}

// Log in the Tessitura client with the given authentication info and cache the login data
func (tq *TqConfig) Login(a auth.Auth) error {

	// Cache the login data
	if basicAuth, err := a.BasicAuth(); err != nil {
		tq.Log.Error(err.Error())
		return err
	} else {
		tq.basicAuth = basicAuth
	}

	host := append(strings.SplitN(a.Hostname, "/", 2), "")
	ignoreCerts, _ := httptransport.TLSClient(httptransport.TLSClientOptions{
		InsecureSkipVerify: true,
	})
	transport := httptransport.NewWithClient(host[0], host[1], []string{"https"}, ignoreCerts)
	transport.DefaultAuthentication = tq.basicAuth
	tq.TessituraServiceWeb = client.New(transport, nil)

	return nil
}

// Generic for doing operations (get/put/post), parallelizing calls to DoOne as needed
// Returns the result, either a single json map if there is only one operation to do,
// or an array of json maps if there are multiple operations.
// Returns the last error
func Do[P any, R any, O any, F func(*P, ...O) (*R, error)](
	tq TqConfig, function F, query []byte,
) ([]byte, error) {
	tq.Log.Info(fmt.Sprint("calling swagger function: ",
		run.FuncForPC(reflect.ValueOf(function).Pointer()).Name()))
	queries := new([]json.RawMessage)
	err := json.Unmarshal(query, queries)
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		tq.Log.Info("query is not an array, so calling API endpoint once")
		// it's not an array... so just call DoOne
		return DoOne(tq, function, query)
	} else if err == nil {
		tq.Log.Info("query is an array, so calling API endpoint multiple times")
		// loop over queries and call DoOne for each
		// TODO: Parallelize this!
		out := make([]json.RawMessage, len(*queries))
		errs := make([]error, 0, len(*queries))
		for i, q := range *queries {
			out[i], err = DoOne(tq, function, q)
			if err != nil {
				errs = append(errs, err)
			}
		}
		res, _ := json.Marshal(out)
		if len(errs) > 0 {
			err = errs[len(errs)-1]
		}
		return res, err
	}
	return nil, err // json.Unmarshal error
}

// Generic for doing operations (get/put/post), parallelizing calls to DoOne as needed
// The type parameters allow it to work for any swagger-defined function.
// It unmarshals `query` into the appropriate data structure for `function` and marshals
// the result back to json.
// If the initial unmarshal fails, it tries again with a depth-first traversal of the
// data structure, essentially trying to recast a flat query into the nested structure
// required.
func DoOne[P any, R any, O any, F func(*P, ...O) (*R, error)](
	tq TqConfig, function F, query []byte,
) ([]byte, error) {

	params := new(P)
	remainder, err := unmarshallStructWithRemainder(query, params)

	// Is there a better way to determine failure?
	if len(structFields(*params)) == 1 && len(mapFields(remainder)) > 0 {
		remainder, err = unmarshallNestedStructWithRemainder(query, params)
	}

	if tq.verbose {
		tq.Log.Info("structFields", "fields", fmt.Sprint(structFields(*params)))
		tq.Log.Info("mapFields", "fields", fmt.Sprint(mapFields(remainder)))
	}
	if len(structFields(*params)) == 0 {
		err = errors.Join(err, fmt.Errorf("query could not be parsed"))
	}
	if tq.dryRun || err != nil {
		return nil, err
	}

	// Call the function
	obj, err := function(params)

	if err != nil {
		return nil, err
	} else {
		// Marshall the json response
		if payload := reflect.ValueOf(obj).Elem().FieldByName("Payload"); !payload.IsZero() {
			return json.Marshal(payload.Interface())
		} else {
			return json.Marshal(obj)
		}
	}
}

// Unmarshall into a struct and return the remainder as a map
// Errors if P is not a struct type
func unmarshallStructWithRemainder(query []byte, params any) (res map[string]any, err error) {
	if reflect.TypeOf(params).Kind() != reflect.Pointer {
		return nil, fmt.Errorf("params must be pointer to struct, got %v", reflect.TypeOf(params).Kind())
	} else if reflect.TypeOf(params).Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("params must be pointer to struct, got pointer to %v", reflect.TypeOf(params).Elem().Kind())
	}

	// Unmarshal the query into the given parameter structure
	_err := json.Unmarshal(query, params)
	err = errors.Join(err, _err)

	// Get all the keys of the query for comparison
	_err = json.Unmarshal(query, &res)
	err = errors.Join(err, _err)

	// Remove keys that are in the struct already
	typ := reflect.TypeOf(params).Elem()
	for key := range res {
		for i := 0; i < typ.NumField(); i++ {
			// go's JSON is case insensitive and so should we
			if strings.EqualFold(key, typ.Field(i).Name) {
				delete(res, key)
			}
		}
	}

	return
}

// Unmarshall into a nested struct from a flat query and return the remainder as a map
// Errors if P is not a struct type
func unmarshallNestedStructWithRemainder(query []byte, params any) (res map[string]any, err error) {
	if reflect.TypeOf(params).Kind() != reflect.Pointer {
		return nil, fmt.Errorf("params must be pointer to struct, got %v", reflect.TypeOf(params).Kind())
	} else if reflect.TypeOf(params).Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("params must be pointer to struct, got pointer to %v", reflect.TypeOf(params).Elem().Kind())
	}

	// First unmarshal the given struct so that those fields get mapped if possible
	res, err = unmarshallStructWithRemainder(query, params)
	if err != nil {
		return
	}
	query, _ = json.Marshal(res)

	v := reflect.ValueOf(params).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Struct && v.Field(i).IsZero() {
			// Recurse if there's a struct field
			res, err = unmarshallNestedStructWithRemainder(query, v.Field(i).Addr().Interface())
			if err != nil {
				return
			}
			// Update the query with only unmatched fields
			query, _ = json.Marshal(res)
		}
	}

	return
}

// Return the field names from struct `s` that are not zero/nil
func structFields(s any) []string {
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}
	usedFields := make([]string, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		// Get the fields actually assigned in the params struct
		if !val.Field(i).IsZero() {
			usedFields = append(usedFields, typ.Field(i).Name)
		}
	}
	return usedFields
}

// Return the field names from map `m`
func mapFields(m map[string]any) []string {
	fields := make([]string, 0, len(m))
	for field := range m {
		fields = append(fields, field)
	}
	return fields
}
