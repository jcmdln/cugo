// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package sha256sum

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

// Sha256sum ...
func (opts *Options) Sha256sum(operands []string) error {
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
