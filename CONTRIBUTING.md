This file is a work-in-progress.

This project was inspired by Rob Landley's project,
[toybox](https://github.com/landley/toybox), and uses the same reference
manuals cited on http://landley.net/toybox/status.html with the
exception being that `cugo` _also_ cares about Unix/BSD compatibility.
When reviewing Unix/BSD reference manuals,
[man.openbsd.org](https://man.openbsd.org/) is the primary source.

Unless you really want something like `cugo` for a specific reason, I
highly suggest contributing to
[toybox](https://github.com/landley/toybox) as it's a much more serious
project than the drivel I'm writing.


## Design Methodology
`cugo` makes use of [Go modules](https://github.com/golang/go/wiki/Modules)
which may or may not be a good idea, but I'm trying to stay ahead of
major shifts in the Go ecosystem. This requires Go 1.11 or later unless
you want to juggle a local patch to modify the behavior or something.

Each utility resides in it's own namespace within `./src`, so that they
may be imported individually into other Go projects. This avoids issues
with Go's default compiler not having link-time optimization, and allows
external projects to ship with only the desired utilities.

See [README](cmd/README.md) for information on the command line
interface provided by `cugo`.

See [README](src/README.md) for information on how to import a single
utility.

See [README](lib/README.md) for information on the libraries written for
`cugo`.


### Utilities
Because `cugo` is a 0-clause BSD licensed project, like `toybox`, the
code for each utility must not be a derivative work of an exiting
utility unless permitted by the license or it's author.

All utilities included in `cugo` must adhere to the following core
tenants:

* Simplicity
* Security
* Portability

#### Simplicity
Be concise and avoid obfuscation.

Any utility included in cugo should be simply correct. The contained
implementations should be written in a concise and readable form without
arbitrarily obfuscating the work being done.

#### Security
The only package used that is not part of the Go standard library is
[flagger](https://github.com/hlfstr/flagger), which is used for the
command line interface within `./cmd`. No other external dependencies
are permitted. This is because I have more faith in the Go maintainers
and contributors to have a proper security review process, publish any
CVE's, and notify their users compared to third-party sources.

#### Portability
While the aim of `cugo` is to be standards compliant, the question
remains of _which_ standards it should comply to. The primary platform
for consideration is Linux, though POSIX specifications take precedence
over GNU-isms. The derivation from standards is not unique to a single
Unix-like variant, as over time each ecosystem has diverged to their own
liking. Despite the implied tone, this is not an insult, I'm simply
upset that I'm trying to wrangle some generality between the standards.

There are many features included within some standards that are not in
others. In such situations, `cugo` will provide whatever seems the most
sane at the time of writing. In some cases, features will be outright
ignored.

An example of a feature that would not be included is GNU's `rm`, which
has `--interactive=INTERVAL` for specifying how many times a user must
say `yes` or `no` before it then proceeds with the rest of the provided
input. These sorts of mechanisms should not be included without a _damn_
good reason.
