// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chmod

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	ex "github.com/jcmdln/cugo/lib/exists"
)

func (opt *Options) Chmod(operands []string) {
	var (
		modeVal uint64
		mode    os.FileMode
		err     error
	)

	if modeVal, err = strconv.ParseUint(operands[0], 8, 32); err != nil {
		fmt.Printf("cugo: %s\n", err)
		os.Exit(1)
	}
	mode = os.FileMode(modeVal)

	for _, operand := range operands[1:] {
		if ex.Exists(operand) {
			if opt.Recursive {
				filepath.Walk(operand, func(t string, info os.FileInfo, err error) error {
					if err = os.Chmod(t, mode); err != nil {
						fmt.Printf("cugo: %s\n", err)
						os.Exit(1)
					}

					return nil
				})
			} else {
				if err = os.Chmod(operand, mode); err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				}
			}
		}
	}

	os.Exit(0)
}
