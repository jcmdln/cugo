// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// compute and check SHA512 message digest
//
// SYNOPSIS
//
//     sha512sum [-bct] [FILE ...]
//
// DESCRIPTION
//
// sha512sum prints or checks SHA512 checksums. If no file is given in the
// operands, read standard input.
//
// SEE ALSO
//
//     tbd
//
// REFERENCES
//
//     https://linux.die.net/man/1/sha512sum
package sha512sum

import (
	"crypto/sha512"
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

// Sha512sum ...
func Sha512sum(args []string) {
	for _, file := range args {
		if contents, err := ioutil.ReadFile(file); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		} else {
			data = []byte(contents)
			fmt.Printf("% x\n", sha512.Sum512(data))
		}
	}
}
