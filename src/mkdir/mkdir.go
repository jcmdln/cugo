// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package mkdir

import (
	"fmt"
	"os"
)

func (opt *Options) Mkdir(operands []string) ([]string, error) {
	var (
		dir  string
		mode os.FileMode
		err  error
	)

	mode = os.FileMode(uint32(opt.Mode))

	for _, dir = range operands {
		if opt.Parents {
			if err = os.MkdirAll(dir, mode); err != nil {
				return operands, err
			}
		} else if err = os.Mkdir(dir, mode); err != nil {
			return operands, err
		}

		if opt.Verbose {
			_, err = fmt.Printf("cugo: mkdir: Created %s\n", dir)
			if err != nil {
				return operands, err
			}
		}
	}

	return operands, nil
}
