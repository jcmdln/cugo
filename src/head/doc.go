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

// display the first few lines of files.
//
// SYNOPSIS
//
//     head [-n NUMBER] [FILE ...]
//
// DESCRIPTION
//
// head copies the first NUMBER of lines of each specified FILE to
// stdout.  If no files are named, head copies lines from stdin.  If
// NUMBER is omitted, it defaults to 10.
//
// The options are as follows:
//
//     -n        Copy the first NUMBER of lines of each file
//
// SEE ALSO
//
//     Not available.
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/head.html
//
package head
