// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// remove files and directories
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
//     http://man.openbsd.org/rm
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/rm.html
package rm
