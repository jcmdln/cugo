// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sha224sum

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

// Sha224sum ...
func (opt *Options) Sha224sum(operands []string) {
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
		fmt.Printf("%x  %s\n", sha256.Sum224(data), operand)
	}

	os.Exit(0)
}
