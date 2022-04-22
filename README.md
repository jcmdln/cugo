`cugo` implements Unix/Linux core utilities as a multi-call binary. Each utility
is written from scratch using the reference manuals and the Go standard library.
Use of external modules is intentionally rare and currently consists of the
following:

- [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys)

If you want an example of importing a single utility into another project,
please see [cugo-mkdir-example](https://github.com/jcmdln/cugo-mkdir-example)
which implements `mkdir` using Go's `flag` package.

# Usage

## Building from source

```sh
$ git clone git@github.com:jcmdln/cugo.git
$ cd ./cugo
$ go build
```

## Installing as a Go module

```sh
$ go get -u github.com/jcmdln/cugo
```
