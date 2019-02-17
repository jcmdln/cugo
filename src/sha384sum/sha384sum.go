// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sha384sum

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

// Sha384sum ...
func (opt *Options) Sha384sum(operands []string) {
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
		fmt.Printf("%x  %s\n", sha512.Sum384(data), operand)
	}

	os.Exit(0)
}
