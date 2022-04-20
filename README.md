`cugo` implements Unix/Linux core utilities as a multi-call binary. Each
utility is written from scratch using the reference manuals and the Go
standard library. Use of external modules is intentionally rare and currently
consists of the following:

* [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys)


About
==========
Please visit [godoc](https://godoc.org/github.com/jcmdln/cugo) for the
latest documentation, or the `doc.go` files within _most_ directories.

* See [cmd/README](cmd/README.md) for information on the command line
interface provided by `cugo`
* See [src/README](src/README.md) for information on how to import a
single utility
* See [lib/README](lib/README.md) for information on the libraries
written for `cugo`


Usage
==========
This repository may have a somewhat odd structure, though it has made
the process of building `cugo` or importing an individual utility into
external projects trivial. No `Makefile`'s or other tools needed.

If you want an example of importing a single utility into another
project, please see
[cugo-mkdir-example](https://github.com/jcmdln/cugo-mkdir-example)
which implements `mkdir` using Go's `flag` package.

Building from source
---
```sh
$ git clone git@github.com:jcmdln/cugo.git
$ cd ./cugo
$ go build
```

Installing as a Go module
---
```sh
$ go get -u github.com/jcmdln/cugo
```

Contributing
===
If you would like to contribute a utility, improve documentation, or
write tests please see the [CONTRIBUTING](CONTRIBUTING.md) file for more
information.
