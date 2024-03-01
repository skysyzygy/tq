# tq - json 🔀 tessitura 

Wrapper around the Tessitura API that reads JSON and outputs a series of API calls to Tessitura. Internally handles authentication, session creation and closure, and batch/concurrent processing so that endpoints can focus on the data and not the intricacies of the API. 

tq is basically a high-level API for common tasks in Tessi...

## usage: 
`tq [options] <verb> <object> (<json> | -f <json-file>)`

### options:
* --dry-run, -n
* --verbose, -v
* --file, -f  
  JSON file to load; if none is specified, reads JSON from stdin

### verbs:  
*verbs can be called by any distinguishable prefix, i.e. "auth", "au", "a" are all abbreviations for "authenticate"*
* authenticate  
  Authenticate with the API server. Credentials are cached (somehow?!) so authentication is only be required once. (Does Tessi do tokens and refreshes? 🤔)
* create, update, retrieve  
  CRU(no D for now to be safe?)
* help <verb>  
  provide documentation for the particular verb

### objects:
* customer
* order

*and in the future?!?*...
* address
* email
* CSI
* ???

## thoughts/questions
```I'm thinking that each verb/object combo could be its own function so the main routine just reads in JSON and dispatches to create.customer, etc.```

```Would be nice if somehow 'help' reads function documentation so that it doesn't have to be its own set of functions.```

```What is the output? Could be a JSON of the responses from Tessi but maybe we just want a success/error output? Or a combination of the two as a JSON?```