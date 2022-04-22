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
