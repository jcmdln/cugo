Cugo is a multi-call binary written in Go which implements core utilities
for various Unix/Linux standards.

```
$ go run main.go -h
cugo: Core Utilities in multi-call Go binary

Usage:
  cugo [flags]
  cugo [command]

Available Commands:
  help        Help about any command
  ls          List files and directories
  mkdir       Create directories
  rm          Remove directory entries
  sleep       delay for a specified amount of time
  touch       Change file access and modification times
  yes         Repeatedly output specified string(s), or 'y'

Flags:
  -h, --help   help for cugo

Use "cugo [command] --help" for more information about a command.
```


## Why?
Cugo was inspired by Rob Landley's toybox though I wanted to play around
with a much simpler build system, and Go allows this without having any
formally defined build system. Go's standard library is feature complete
enough to make many utilities trivial to implement, sometimes with little
more than meeting the conditions required from reference documentation.


## Design Methodology
Any utility included in cugo should be simply correct. The contained
implementations should be written in a concise and readable form without
arbitrarily obfuscating the work being done.

While the aim of Cugo is to be standards compliant, there are many
features included within some standards that are not in others. In such
situations, a happy medium between available features will be chosen. An
example of this would be `GNU rm` which  has `--interactive=INTERVAL` for specifying how many times a user must say `yes` or `no` before it then
proceeds with the rest of the provided input. These sorts of built-in
script-ish mechanisms will not be included in any utility unless it is
vital to functionality, even at the sake of violating specifications.
