Cugo is a multi-call binary written in Go which implements core utilities
from various Unix/Linux standards.

```
$ go run cugo.go -h
Core Utilities in multi-call Go binary

Usage:
  cugo [flags]
  cugo [command]

Available Commands:
  basename    Return non-directory portion of a pathname
  chmod       Change file mode bits
  count       Count the number of elements of an array
  false       Return false value
  help        Help about any command
  hostname    Return the host name reported by the kernel
  ls          List files and directories
  mkdir       Create directories
  pwd         Return working directory name
  rm          Remove files and directories
  rmdir       Remove directories
  sleep       Delay for a specified amount of time
  touch       Change file access and modification times
  true        Return true value
  whoami      Return current user
  yes         Repeatedly output specified string(s), or 'y'

Flags:
  -h, --help   help for cugo

Use "cugo [command] --help" for more information about a command.
```


## Why?
Cugo was inspired by Rob Landley's toybox though I wanted to play around
with a much simpler build system. Go is able to facillitate a multi-call
binary without requiring a formally defined build system. Go's standard
library is feature complete enough to make many utilities trivial to
implement, sometimes with little more than meeting the conditions
required by reference documentation.


## Design Methodology
Any utility included in cugo should be simply correct. The contained
implementations should be written in a concise and readable form without
arbitrarily obfuscating the work being done.

While the aim of Cugo is to be standards compliant, there are many
features included within some standards that are not in others. In such
situations, a happy medium between reference manuals will be chosen. An
example of this would be `GNU rm` which  has `--interactive=INTERVAL` for
specifying how many times a user must say `yes` or `no` before it then
proceeds with the rest of the provided input. These sorts of built-in
script-ish mechanisms will not be included in any utility unless it is
vital to functionality, even at the sake of violating specifications.

Suitable `man` page sources:
* https://man.openbsd.org/
* http://landley.net/toybox/status.html


## Contributing
The only package used that is not part of the Go standard library is
Cobra and it should remain this way. Because this is a 0-clause BSD
project the code for each utility must not be a derivative work of an
exiting utility.

Feel free to submit a pull request, file an issue, or join the
discussion at #cugo on FreeNode.
