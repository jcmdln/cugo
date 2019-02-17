[Homepage](https://cugo.io) | [GitHub](https://github.com/jcmdln/cugo)


`cugo` is a prototypical project to implement Unix/Linux core utilities
in the form of a multi-call binary, primarily using the Go standard 
library. Each utility is written from scratch using only the reference 
manuals for the target utility and the Go standard library. Third-party
code use by `cugo` is intentionally rare, currently relying on the
following:

* [hlfstr/flagger](https://github.com/hlfstr/flagger) for the optional
command line interface that resides within the `cmd` folder.
* [x/sys](https://godoc.org/golang.org/x/sys) for various system calls.


## About
Please visit [godoc](https://godoc.org/github.com/jcmdln/cugo) for the
latest documentation, or the `doc.go` files within _most_ directories.

* See [cmd/README](cmd/README.md) for information on the command line
interface provided by `cugo`.
* See [src/README](src/README.md) for information on how to import a 
single utility.
* See [lib/README](lib/README.md) for information on the libraries 
written for `cugo`.


## Usage
This repository may have a somewhat odd structure, though it has made
the process of building `cugo` or importing an individual utility into
external projects trivial. No `Makefile`'s or other tools needed.

### Building

	$ git clone git@github.com:jcmdln/cugo.git
	$ cd ./cugo
	$ go run cugo.go
	# OR 
	$ go build

### Installing

	$ go get -u github.com/jcmdln/cugo

### Importing
If you want an example of importing a single utility into another 
project, please see [cugo-mkdir-example](https://github.com/jcmdln/cugo-mkdir-example)
which implements `mkdir` using Go's `flag` package.


## Contributing
If you would like to contribute a utility, improve documentation, or
write tests please see the [CONTRIBUTING](CONTRIBUTING.md) file for more
information.
