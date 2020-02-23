// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

// copy the last part of a file
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
