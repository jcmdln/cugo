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

// +build testing

// list directory contents.
//
// SYNOPSIS
//
//     ls [-aR] [FILE ...]
//
// DESCRIPTION
//
// ls receives operands and may display the name (or other information)
// of the provided file or directory.  For each named directory, ls
// displays the names of files contained within that directory as well
// as any requested, associated information.
//
// The options are as follows:
//
//     -A        Show all entries other than '.' and '..'.
//
//     -a        Show all entries including those starting with '.'.
//
//     -F        Add '/' after each directory, '*' after executables,
//               '@' after symlinks, '=' after sockets, and '|' after FIFO.
//
//     -H        Follow symlinks.
//
//     -h        Use human-readable units.
//
//     -L        Evaluate symlinks as physical references, but use the
//               name of the symlink rather than the referenced file.
//
//     -l        Show entries in a list using long format, one entry
//               per line.
//
//     -P        Add '/' after each directory name.
//
//     -R        Recursively list directories and their entries.
//
//     -r        Reverse the sorting order.
//
//     -S        Sort by size, largest file first.
//
//
// SEE ALSO
//
//     https://golang.org/pkg/io/ioutil/#ReadDir
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/ls.html
//
package ls
