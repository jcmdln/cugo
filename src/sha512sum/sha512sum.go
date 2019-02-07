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

type Options struct {
	Binary bool
	Check  bool
	Text   bool
}

type Opts func(*Options)

func Binary(binary bool) Opts {
	return func(opts *Options) {
		opts.Binary = binary
	}
}

func Check(check bool) Opts {
	return func(opts *Options) {
		opts.Check = check
	}
}

func Text(text bool) Opts {
	return func(opts *Options) {
		opts.Text = text
	}
}

// Sha512sum ...
func (opt *Options) Sha512sum(operands []string) {
	var (
		operand  string
		contents []byte
		data     []byte
		err      error
	)

	for _, operand = range operands {
		if contents, err = ioutil.ReadFile(operand); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}

		data = []byte(contents)
		fmt.Printf("%x  %s\n", sha512.Sum512(data), operand)
	}

	os.Exit(0)
}
