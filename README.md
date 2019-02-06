[Homepage](https://cugo.io) | [GitHub](https://github.com/jcmdln/cugo)


NOTE: `cugo` is currently pre-alpha software that has no tests and
little documentation. Please review this repository in its entirety
before usage, as you may experience undesired or undocumented behavior!


## About
`cugo` is a pet project to re-implement Unix/Linux core utilities in the
form of a multi-call binary. Each utility is written from scratch using
only the reference manuals for the target utility, and the Go standard
library.

Please visit [godoc](https://godoc.org/github.com/jcmdln/cugo) for the
latest documentation.

See [README](cmd/README.md) for information on the command line
interface provided by `cugo`.

See [README](src/README.md) for information on how to import a single
utility.

See [README](lib/README.md) for information on the libraries written for
`cugo`.


## Building and Installing
First, check out the command line and the available commands:

	go run cugo.go -h

You may build `cugo` using the following:

	go build cugo.co

If you want a smaller binary and don't care about debugging, try the
following:

	go build -ldflags="-w -s" cugo.go

Once compiled, the utilities run fairly quick and shouldn't be painful
compared to the performance of GNU bin-utils or other core utilities. I
am of the opinion that "more faster is more better" but `cugo` isn't at
the point where investigating avenues for optimization is a high
priority. If you have some suggestions, reach out any way you see fit.


## Contributing
Please see the [CONTRIBUTING](CONTRIBUTING.md) file for more information.
