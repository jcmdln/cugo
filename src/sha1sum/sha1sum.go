// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// compute and check SHA1 message digest
//
// SYNOPSIS
//
//     sha1sum [-bct] [FILE ...]
//
// DESCRIPTION
//
// sha1sum prints or checks SHA1 checksums. If no file is given in the
// operands, read standard input.
//
// SEE ALSO
//
//     tbd
//
// REFERENCES
//
//     https://linux.die.net/man/1/sha1sum
package sha1sum

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	// Binary is a boolean that, when true, enables reading in binary
	// mode.
	Binary bool
	// Check is a boolean that, when true, reads sums from FILEs and
	// checks them.
	Check bool
	// Text is a bool that, when true, enables reading in binary mode,
	// though this is the default behavior.
	Text bool

	data []byte
)

// Sha1sum ...
func Sha1sum(args []string) {
	for _, file := range args {
		if contents, err := ioutil.ReadFile(file); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		} else {
			data = []byte(contents)
			fmt.Printf("% x\n", sha1.Sum(data))
		}
	}
}
