// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// compute and check SHA256 message digest
//
// SYNOPSIS
//
//     sha256sum [-bct] [FILE ...]
//
// DESCRIPTION
//
// sha256sum prints or checks SHA256 checksums. If no file is given in the
// operands, read standard input.
//
// SEE ALSO
//
//     tbd
//
// REFERENCES
//
//     https://linux.die.net/man/1/sha256sum
package sha256sum

import (
	"crypto/sha256"
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

// Sha256sum ...
func Sha256sum(args []string) {
	for _, file := range args {
		if contents, err := ioutil.ReadFile(file); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		} else {
			data = []byte(contents)
			fmt.Printf("% x\n", sha256.Sum256(data))
		}
	}
}
