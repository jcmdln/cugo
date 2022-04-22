// doc.go
//
// ISC License
//
// Copyright (c) 2022 Johnathan C. Maudlin <jcmdln@gmail.com>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE
// OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

//go:build testing
// +build testing

// copy the last part of a file.
//
// SYNOPSIS
//
//     tail [-f] [-c NUMBER | -n NUMBER] [FILE ...]
//
// DESCRIPTION
//
// tail copies each input file to stdout beginning at a designated
// location.
//
// Copying begins at the point in the file indicated by the NUMBER
// passed to the -c and -n options.
//
// The options are as follows:
//
//     -c        The location is NUMBER bytes.
//
//     -f        When reaching EOF, wait for additional data to be
//               appended to the input.  If the file is replaced, tail
//               will re-open the file and continue.
//
//     -n        The location is NUMBER lines.
//
// SEE ALSO
//
//     Not available.
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/tail.html
//
package tail
