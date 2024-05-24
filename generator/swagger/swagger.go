// WARNING! This will update all of the client API definitions and will take a LONG time.

// to make a new swagger client and model definition, place the updated swagger json in the generator directory and run `go generate` in this one.
// you then may need to run go generate in the parent directory to update the tq command files.

//go:generate go install github.com/go-swagger/go-swagger/cmd/swagger@latest

// fix up the json with some default error responses and make pass swagger spec
//go:generate sh -c "jq -f tessitura.jq ../*.json > tessitura.json"

//go:generate sh -c "$(go env GOPATH)/bin/swagger generate client -f tessitura.json -t ../../"

package swagger
