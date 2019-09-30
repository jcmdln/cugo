// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// change file mode bits.
//
// SYNOPSIS
//
//     chmod [-R] MODE FILE ...
//
// DESCRIPTION
//
// Chmod changes the file mode bits of provided files as specified by
// the MODE operand.  The MODE of a file determines the permissions and
// attributes of the file in question.  Currently the only supported
// MODE operand format requires using octal numbers from 0 to 7.
//
// The options are as follows:
//
//     -R        Change files and directories recursively.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#Chmod
//
// REFERENCES
//
// 	   http://man.openbsd.org/chmod
// 	   http://pubs.opengroup.org/onlinepubs/9699919799/utilities/chmod.html
package chmod
