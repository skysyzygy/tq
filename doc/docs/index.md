### > **tq** is a command-line interface for Tessitura. 🚀

**tq** is a wrapper around the Tessitura API that reads JSON-formatted data and executes a series of API calls to Tessitura. It internally handles authentication, session creation and closure, and batch/concurrent processing so that humans like you (or bots or scripts) can focus on the data and not the intricacies of the Tessitura API.                                                       
                     
# 🏗️ installation

## from binary

Download the latest release from the [releases page!](https://github.com/skysyzygy/tq/releases/) 

## from source

The only prerequisite to using **tq** is [installing go](https://go.dev/doc/install).

Then clone this repository and build:
```shell
git clone github.com/skysyzygy/tq
cd tq
go build -o bin/tq .
```
The build command will create an executable file `tq` or `tq.exe` in the `bin` project directory.

# 🪪 authentication

To authenticate with the API server you need to select at least one authentication method. 
```shell
tq auth add --host hostname --user username --group usergroup --location location
# Password: ******

tq auth sel --host hostname --user username --group usergroup --location location
```

# 🍳 recipes

**get constituent info**
```shell
tq get constituents <<< '{"constituentid": "12345"}'
```
**update a constituent address**
```shell
tq update addresses <<< '{"addressid": "12345", "street1": "123 New Street"}'
```
**add a plan step**
```shell
tq create steps <<< '{"plan": {"Id": 12345}, "type": {"Id": 1}, "Description": "New step!", "Notes": "Created by tq :)"}'
# or using flattened syntax
tq create steps <<< '{"plan.Id": 12345, "type.Id": 1, "Description": "New step!", "Notes": "Created by tq :)"}'
```

# 🛠️ usage

```shell 
tq [flags] [verb] [object]
```

## flags:
*  **-c, --compact** compact instead of indented output
*  **-n, --dryrun** don't actually do anything, just show what would have happened
*  **-f, --file** input file to read (default is to read from stdin)
*  **--highlight** render json with syntax highlighting; default is to use highlighting when output is to terminal
*  **-i, --in** input format (csv or json; default is json); csv implies --inflat
*  **--inflat** use input flattened by JSONPath dot notation. Combining this with --help will show the flattened format
*  **-l, --log** log file to write to (default is no log)
*  **--no-highlight** render json without syntax highlighting; default is to use highlighting when output is to terminal
*  **-o, --out** output format (csv or json; default is json); csv implies --outflat
*  **--outflat** use output flattened by JSONPath dot notation
*  **-v, --verbose** turns on additional diagnostic output


#### configuration file:
A yaml configuration file `.tq` placed in your home directory can be used to set defaults for these flags; it is also used to save the current authentication method. See `tq auth select --help` for more information. 


### verbs:
*  **authenticate** : Authenticate with the Tessitura API
*  **create** :       Create entities in Tessitura
*  **get** :          Retrieve entities from Tessitura
*  **update** :       Update entities in Tessitura
*  **completion** :   Generate the autocompletion script for the specified shell
*  **help** :         Help about any command

## queries:
Queries are simply JSON objects and can be batched by combining multiple query objects into a single JSON array, e.g. 

```json
[{"ID":123}, {"ID":124}, {"ID":125}, {}, {}, ]
```
Query object details are detailed in the help for each command.

**Queries can be sent to `tq` by:**

1. writing them to a file, e.g. a file like this:

> query.json
```json
{"CustomerId":"12345"}
```
```shell
tq -f query.json get constituents
# ...or
tq get constituents < query.json
```

2. By piping them on the command line to `tq` directly:
```shell
echo {"CustomerId":"12345"} | tq get constituents
```

3. By using a here-string:
```shell
# bash
tq get constituents <<< '{"CustomerId":"12345"}'
```
```shell
# powershell
'{\"CustomerId\":\"12345\"}' | tq get constituents
```

4. Or for longer queries, using a here-doc!
```shell
# bash
tq get constituents <<EOF 
[
      {"CustomerId":"12345"},
      {"CustomerId":"12346"},
      {"CustomerId":"12347"},
      {"CustomerId":"12348"}
]
EOF
```
```shell
# powershell
@'
[
      {"CustomerId":"12345"},
      {"CustomerId":"12346"},
      {"CustomerId":"12347"},
      {"CustomerId":"12348"}
]
'@ | tq get constituents
```