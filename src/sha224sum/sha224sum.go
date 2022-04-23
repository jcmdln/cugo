// SPDX-License-Identifier: ISC

package sha224sum

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

type Option struct {
	Binary bool
	Check  bool
	Text   bool
}

// Sha224sum ...
func (opt *Option) Sha224sum(operands []string) error {
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
		if _, err = fmt.Printf("%x  -\n", sha256.Sum224(data)); err != nil {
			return err
		}
	} else {
		for _, operand = range operands {
			if contents, err = ioutil.ReadFile(operand); err != nil {
				return err
			}

			data = []byte(contents)
			if _, err = fmt.Printf("%x  %s\n", sha256.Sum224(data), operand); err != nil {
				return err
			}
		}
	}

	return nil
}
