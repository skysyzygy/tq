package main

import (
	"encoding/json"
	"reflect"
	"slices"
	"testing"

	"github.com/skysyzygy/tq/client"
	"github.com/skysyzygy/tq/client/g_e_t"
	"github.com/skysyzygy/tq/client/p_o_s_t"
	"github.com/skysyzygy/tq/client/p_u_t"
	"github.com/stretchr/testify/assert"
)

func Test_Describe(t *testing.T) {

	short, long := Describe("ConstituentsGet")
	assert.Equal(t, "Constituents", short)
	assert.Equal(t, "Get the details of a Constituent using id.\n ConstituentId(*): string\t", long)

	short, long = Describe("ConstituentsUpdate")
	assert.Equal(t, "Constituents", short)
	assert.Equal(t, "Update a constituent. Only the information about constituent can be updated. If addresses, electronicAddresses, salutations or phones data are passed, they will be ignored.\n ConstituentId(*): string\t\n Constituent(*): object\t", long)

	short, long = Describe("ConstituentsCreateConstituent")
	assert.Equal(t, "Constituents", short)
	assert.Equal(t, "Create a new constituent with addresses, electronicAddresses, salutations and phones.\n Constituent(*): object\t", long)

	short, long = Describe("ElectronicAddressesGet")
	assert.Equal(t, "ElectronicAddresses", short)
	assert.Equal(t, "Get details of an electronic address.\n ElectronicAddressId(*): string\t", long)

	short, long = Describe("ElectronicAddressesUpdate")
	assert.Equal(t, "ElectronicAddresses", short)
	assert.Equal(t, "Update an electronic address.\n ElectronicAddressId(*): string\t\n ElectronicAddress(*): object\t", long)

	short, long = Describe("ElectronicAddressesCreate")
	assert.Equal(t, "ElectronicAddresses", short)
	assert.Equal(t, "Create a new electronic address.\n ElectronicAddress(*): object\t", long)

	short, long = Describe("AssociationTypesGet")
	assert.Equal(t, "AssociationTypes", short)
	assert.Equal(t, "Get the details of an association type by id. To get the resource only if the user has write/edit access, pass filter=\"writeonly\". To get the resources in maintenance mode(ignore Control Groups), pass maintenanceMode=\"true\". Maintenance mode requires users to have access to the reference tables. (Specified in TX_REFTABLE_USERGROUP).\n Id(*): string\t\n Filter: string\tFilter by user access (default: readwrite)\n MaintenanceMode: string\tIgnore control grouping (default: false)", long)

}

// Test that instantiateStructType recurses into a struct, navigating pointers, and initializes
// strings with "string" and everything else with zero type
func Test_instantiateStructType(t *testing.T) {
	type test struct {
		I  int
		F  float32
		B  bool
		S  string
		SP *string
	}
	type nest struct {
		P *test
		O test
	}
	pointer := new(nest)
	object := *pointer
	j, _ := json.Marshal(instantiateStructType(reflect.TypeOf(object), 1))
	assert.Equal(t,
		`{"P":{"I":0,"F":0,"B":false,"S":"string","SP":"string"},"O":{"I":0,"F":0,"B":false,"S":"string","SP":"string"}}`,
		string(j))

	j, _ = json.Marshal(instantiateStructType(reflect.TypeOf(pointer), 1))
	assert.Equal(t,
		`{"P":{"I":0,"F":0,"B":false,"S":"string","SP":"string"},"O":{"I":0,"F":0,"B":false,"S":"string","SP":"string"}}`,
		string(j))

	j, _ = json.Marshal(instantiateStructType(reflect.TypeOf(pointer), 0))
	assert.Equal(t,
		`{"P":null,"O":{"I":0,"F":0,"B":false,"S":"","SP":null}}`,
		string(j))

}

func Test_usage(t *testing.T) {
	assert.Equal(t, `{"ConstituentID":"string"}`,
		string(usage(reflect.ValueOf(g_e_t.ClientService.ConstituentsGet))))
	assert.Equal(t, `{"Constituent":{"CreateLocation":"string","CreatedBy":"string","CreatedDateTime":"0001-01-01T00:00:00.000Z","DisplayName":"string","FirstName":"string","LastActivityDate":"0001-01-01T00:00:00.000Z","LastGiftDate":"0001-01-01T00:00:00.000Z","LastName":"string","LastTicketDate":"0001-01-01T00:00:00.000Z","MiddleName":"string","SortName":"string","UpdatedBy":"string","UpdatedDateTime":"0001-01-01T00:00:00.000Z"},"ConstituentID":"string"}`,
		string(usage(reflect.ValueOf(p_u_t.ClientService.ConstituentsUpdate))))
	assert.Equal(t, `{"Constituent":{"Addresses":[object],"CreateLocation":"string","CreatedBy":"string","CreatedDateTime":"0001-01-01T00:00:00.000Z","DisplayName":"string","ElectronicAddresses":[object],"FirstName":"string","LastActivityDate":"0001-01-01T00:00:00.000Z","LastGiftDate":"0001-01-01T00:00:00.000Z","LastName":"string","LastTicketDate":"0001-01-01T00:00:00.000Z","MiddleName":"string","PhoneNumbers":[object],"Salutations":[object],"SortName":"string","UpdatedBy":"string","UpdatedDateTime":"0001-01-01T00:00:00.000Z"}}`,
		string(usage(reflect.ValueOf(p_o_s_t.ClientService.ConstituentsCreateConstituent))))
}

func Test_newCommandGet(t *testing.T) {
	client := client.New(nil, nil)
	c := reflect.TypeOf(client.Get)
	commands := make([]command, c.NumMethod())
	for i := 0; i < c.NumMethod(); i++ {
		if methodName := c.Method(i).Name; slices.Contains([]string{
			"SetTransport", "CacheInit", "DiagnosticsGetAPIPluginConfiguration"},
			methodName) {
			continue
		}
		commands[i] = newCommand(c.Method(i))
		assert.NotEqual(t, "", commands[i].verb, commands[i])
		assert.NotEqual(t, "", commands[i].thing, commands[i])
	}
}

func Test_newCommandPut(t *testing.T) {
	client := client.New(nil, nil)
	c := reflect.TypeOf(client.Put)
	commands := make([]command, c.NumMethod())
	for i := 0; i < c.NumMethod(); i++ {
		if methodName := c.Method(i).Name; slices.Contains([]string{
			"SetTransport", "DiagnosticsPut"},
			methodName) {
			continue
		}
		commands[i] = newCommand(c.Method(i))
		assert.NotEqual(t, "", commands[i].verb, commands[i])
		assert.NotEqual(t, "", commands[i].thing, commands[i])
	}
}

func Test_newCommandPost(t *testing.T) {
	client := client.New(nil, nil)
	c := reflect.TypeOf(client.Post)
	commands := make([]command, c.NumMethod())
	for i := 0; i < c.NumMethod(); i++ {
		if methodName := c.Method(i).Name; slices.Contains([]string{
			"SetTransport",
			"EmailsSendTickets", "OrdersPrintHTMLTickets"},
			methodName) {
			continue
		}
		commands[i] = newCommand(c.Method(i))
		assert.NotEqual(t, "", commands[i].verb, commands[i])
		assert.NotEqual(t, "", commands[i].thing, commands[i])
	}
}
