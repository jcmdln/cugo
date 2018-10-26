`cugo` provides common Unix core utilities in the form of a multi-call
binary, with a focus on broad support for various Unix and Unix-like
operating systems.


## Design Methodology
Any utility included in cugo should be simply correct. The contained
implementations should be written in a concise and readable form without
arbitrarily obfuscating the work being done.

While the aim of Cugo is to be standards compliant, there are many
features included within some standards that are not in others. In such
situations, a happy medium between reference manuals will be chosen.

Suitable `man` page sources:
* https://man.openbsd.org/
* http://landley.net/toybox/status.html

An example of this would be `GNU rm` which  has `--interactive=INTERVAL`
for specifying how many times a user must say `yes` or `no` before it
then proceeds with the rest of the provided input. These sorts of
awkward mechanisms should not be included without good reason.


## Documentation
Please visit https://godoc.org/github.com/jcmdln/cugo for the latest
usage information. If any information is incorrect it is considered a bug
and I invite you to submit an issue, please and thank you!


## Contributing
The only package used that is not part of the Go standard library is
Cobra and it should remain this way. Because this is a 0-clause BSD
project the code for each utility must not be a derivative work of an
exiting utility.

Feel free to submit a pull request, file an issue, or join the
discussion at #cugo on FreeNode.
