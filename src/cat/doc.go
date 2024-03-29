// SPDX-License-Identifier: ISC

// concatenate and print files.
//
// SYNOPSIS
//
//     cat [-u] [FILE ...]
//
// DESCRIPTION
//
// cat reads files sequentially, writing them to stdout.  The FILE
// operands are processed in the provided order.  If FILE is a single
// dash ('-') or absent, cat reads from stdin.
//
// The options are as follows:
//
//     -u        Unbuffered output
//
// SEE ALSO
//
// tbd
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/cat.html
//     http://harmful.cat-v.org/cat-v/
//
package cat
