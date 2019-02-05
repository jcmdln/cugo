// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// compute and check SHA384 message digest
//
// SYNOPSIS
//
//     sha384sum [-bct] [FILE ...]
//
// DESCRIPTION
//
// sha384sum prints or checks SHA384 checksums. If no file is given in the
// operands, read standard input.
//
// SEE ALSO
//
//     tbd
//
// REFERENCES
//
//     https://linux.die.net/man/1/sha384sum
package sha384sum

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

// Sha384sum ...
func Sha384sum(args []string) {
	for _, file := range args {
		if contents, err := ioutil.ReadFile(file); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		} else {
			data = []byte(contents)
			fmt.Printf("%x  %s\n", sha512.Sum384(data), file)
		}
	}

	os.Exit(0)
}
