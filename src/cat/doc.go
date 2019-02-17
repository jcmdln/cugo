// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// concatenate and print files
//
// SYNOPSIS
//
//     cat [-u] [FILE ...]
//
// DESCRIPTION
//
// cat reads files sequentially, writing them to standard out. the FILE
// operands are processed in the provided order. If FILE is a single
// dash ('-') or absent, cat reads from standard in.
//
// The options are as follows:
//
//     -u        Unbuffered output
//
// SEE ALSO
//
// tbd
//
// REFERENCES
//
//     http://man.openbsd.org/cat
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/cat.html
package cat
