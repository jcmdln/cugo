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

// return non-directory portion of a pathname.
//
// SYNOPSIS
//
//     basename STRING [suffix]
//
// DESCRIPTION
//
// basename deletes any prefix ending with the last slash ("/")
// character present in STRING, and a suffix, if given.  The resulting
// filename is written to the standard output.  A non-existent suffix
// is ignored.
//
// SEE ALSO
//
//     https://golang.org/pkg/path/filepath/#Base
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/basename.html
//
package basename
