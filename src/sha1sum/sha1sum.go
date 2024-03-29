// SPDX-License-Identifier: ISC

package sha1sum

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

type Option struct {
	Binary bool
	Check  bool
	Text   bool
}

// Sha1sum ...
func (opt *Option) Sha1sum(operands []string) error {
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
		if _, err = fmt.Printf("%x  -\n", sha1.Sum(data)); err != nil {
			return err
		}
	} else {
		for _, operand = range operands {
			if contents, err = ioutil.ReadFile(operand); err != nil {
				return err
			}

			data = []byte(contents)
			if _, err = fmt.Printf("%x  %s\n", sha1.Sum(data), operand); err != nil {
				return err
			}
		}
	}

	return nil
}
