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

import (
	"fmt"
	"io"
	"os"
)

var (
	// Unbuffered is a boolean that, when true, enables unbuffered
	// output.
	Unbuffered bool

	operand string
	file    io.Reader
	buffer  []byte
	buff    = make([]byte, 1)
	err     error
)

// Cat ...
func Cat(operands []string) {
	if Unbuffered {
		buffer = buff
	} else {
		buffer = make([]byte, 4096)
	}

	if len(operands) == 0 {
		io.CopyBuffer(os.Stdout, os.Stdin, buffer)
	} else {
		for _, operand = range operands {
			if file, err = os.Open(operand); err != nil {
				fmt.Printf("cugo: %s", err)
				os.Exit(1)
			}

			if _, err = file.Read(buff); err != nil {
				fmt.Printf("cugo: %s", err)
				os.Exit(1)
			}

			io.CopyBuffer(os.Stdout, file, buffer)
		}
	}

	os.Exit(0)
}
