// SPDX-License-Identifier: ISC

package sha256sum

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

// Sha256sum ...
func (opts *Option) Sha256sum(operands []string) error {
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
		if _, err = fmt.Printf("%x  -\n", sha256.Sum256(data)); err != nil {
			return err
		}
	} else {
		for _, operand = range operands {
			if contents, err = ioutil.ReadFile(operand); err != nil {
				return err
			}

			data = []byte(contents)
			if _, err = fmt.Printf("%x  %s\n", sha256.Sum256(data), operand); err != nil {
				return err
			}
		}
	}

	return nil
}
