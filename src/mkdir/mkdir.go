// SPDX-License-Identifier: ISC

package mkdir

import (
	"fmt"
	"os"
)

type Option struct {
	Mode    uint
	Parents bool
	Verbose bool
}

func (opt *Option) Mkdir(operands []string) ([]string, error) {
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
