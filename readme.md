 
![go workflow](https://github.com/skysyzygy/tq/actions/workflows/go.yml/badge.svg)
[![codecov](https://codecov.io/gh/skysyzygy/tq/graph/badge.svg?token=Ov8gpBWhHQ)](https://codecov.io/gh/skysyzygy/tq)

> **tq** is a command-line interface for Tessitura. 🚀

**tq** is a wrapper around the Tessitura API that reads JSON-formatted data and executes a series of API calls to Tessitura. It internally handles authentication and batch/concurrent processing so that humans like you (or bots or scripts) can focus on the data and not the intricacies of the Tessitura API.                                                       

# 📚 documentation

Full documentation can be found at [skysyzygy.xyz/tq](https://skysyzygy.xyz/tq)

# 🏗️ installation

## from binary

Download the latest release from the [releases page!](https://github.com/skysyzygy/tq/releases/) 

## from source

The only prerequisite to building **tq** is [installing go](https://go.dev/doc/install).

Then clone this repository and build:
```shell
git clone github.com/skysyzygy/tq
cd tq
go build -o bin/tq .
```
The build command will create an executable file `tq` or `tq.exe` in the `bin` project directory.
