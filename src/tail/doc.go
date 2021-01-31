// doc.go
//
// Copyright 2021 Johnathan C. Maudlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

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
