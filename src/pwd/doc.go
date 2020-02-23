// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

// return working directory name
//
// SYNOPSIS
//
//     pwd [-LP]
//
// DESCRIPTION
//
// The pwd utility prints the absolute pathname of the current working
// directory to the standard output.
//
// The options are as follows:
//
//     -L        Print the logical (including symlinks) path of the
//               current directory.
//
//     -P        Print the physical path of the current directory.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#Getwd
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/pwd.html
//
package pwd
