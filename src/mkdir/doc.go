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

// create directories.
//
// SYNOPSIS
//
//     mkdir [-p] [-m MODE] DIRECTORY ...
//
// DESCRIPTION
//
// mkdir creates the directories named as arguments in the order
// specified using mode rwxrwxrwx (0777) as modified by the current
// umask.
//
// The options are as follows:
//
//     -m        Set the file permission bits of a new directory to MODE.
//
//     -p        Create missing parent directories as required.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#FileMode
//     https://golang.org/pkg/os/#Mkdir
//     https://golang.org/pkg/os/#MkdirAll
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/mkdir.html
//
package mkdir
