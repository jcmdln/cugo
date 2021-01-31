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

// remove files and directories.
//
// SYNOPSIS
//
//     rm [-dfiRrv] FILE ...
//
// DESCRIPTION
//
// rm attempts to remove non-directory type files specified as arguments
// regardless of the file's permissions.
//
// The options are as follows:
//
//     -d        Remove empty directories.
//
//     -f        Skip prompts and ignore warnings.
//
//     -i        Prompt before each removal.
//
//     -R, -r    Remove directories and their contents recursively.
//
//     -v        Print a message when actions are taken.
//
// It is an error to attempt to remove the root directory or the files
// “.” or “..”. It is forbidden to remove the file “..” merely to avoid
// the antisocial consequences of inadvertently doing something like
// “rm -r .*”.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#Stat
//     https://golang.org/pkg/os/#Remove
//     https://golang.org/pkg/path/filepath/#Walk
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/rm.html
//
package rm
