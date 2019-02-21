// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sha1sum

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

// Sha1sum ...
func (opt *Options) Sha1sum(operands []string) {
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
		fmt.Printf("%x  %s\n", sha1.Sum(data), operand)
	}

	os.Exit(0)
}
