// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// list directory contents
//
// SYNOPSIS
//
//     ls [-aR] [FILE ...]
//
// DESCRIPTION
//
// ls receives operands and may display the name (or other information)
// of the provided file or directory. For each named directory, ls
// displays the names of files contained within that directory as well
// as any requested, associated information.
//
// The options are as follows:
//
//     -a        Include directory entries that start with a dot ('.').
//
//     -R        Recursively list encountered subdirectories.
//
// SEE ALSO
//
//     https://golang.org/pkg/io/ioutil/#ReadDir
//
// REFERENCES
//
//     http://man.openbsd.org/ls
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/ls.html
package ls
