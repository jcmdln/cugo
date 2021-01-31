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

// change file access and modification times.
//
// SYNOPSIS
//
//     touch [-acm] [-d ccyy-mm-ddTHH:MM:SS[.frac][Z]] [-r FILE] FILE ...
//
// DESCRIPTION
//
// touch sets the modification and access times of files to the current
// time of day.  If the file doesn't exist, it is created with default
// permissions.
//
// The options are as follows:
//
//     -a        Change the access time of the file.
//
//     -c        Do not create missing files, or display an error when a
//               file is either missing or not created.
//
//     -d        Change the access and modified times of the file, using
//               ISO8601/RFC3339Nano format.
//
//     -m        Change the modified time of the file.
//
//     -r        Use the access and modified times of the reference file
//               rather than the current date-time.
//
// SEE ALSO
//
//     https://golang.org/pkg/time/#pkg-constants
//     https://golang.org/pkg/os/#Chtimes
//     https://golang.org/pkg/time/#Unix
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/touch.html
//
package touch
