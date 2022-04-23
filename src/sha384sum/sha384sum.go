// SPDX-License-Identifier: ISC

package sha384sum

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

type Option struct {
	Binary bool
	Check  bool
	Text   bool
}

// Sha384sum ...
func (opt *Option) Sha384sum(operands []string) error {
	var (
		contents []byte
		data     []byte
		err      error
		operand  string
	)

	if len(operands) < 1 {
		if contents, err = ioutil.ReadAll(os.Stdin); err != nil {
			return err
		}

		data = []byte(contents)
		if _, err = fmt.Printf("%x  -\n", sha512.Sum384(data)); err != nil {
			return err
		}
	} else {
		for _, operand = range operands {
			if contents, err = ioutil.ReadFile(operand); err != nil {
				return err
			}

			data = []byte(contents)
			if _, err = fmt.Printf("%x  %s\n", sha512.Sum384(data), operand); err != nil {
				return err
			}
		}
	}

	return nil
}
