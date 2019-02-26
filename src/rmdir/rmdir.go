// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rmdir

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	em "github.com/jcmdln/cugo/lib/empty"
	ex "github.com/jcmdln/cugo/lib/exists"
)

func (opt *Options) Rmdir(operands []string) error {
	var (
		operand string
		err     error
	)

	rmdir := func(dir string) error {
		if err = os.Remove(dir); err != nil {
			return err
		}

		if opt.Verbose {
			fmt.Printf("cugo: rm: Removed '%s'\n", dir)
		}
		return nil
	}

	for _, operand = range operands {
		if ex.Exists(operand) != nil {
			err = errors.New("rmdir: " + operand + ": no such file or directory")
			return err
		}

		if em.Empty(operand) {
			if err = rmdir(operand); err != nil {
				return err
			}
		} else if opt.Parents {
			for !em.Empty(operand) {
				if err = filepath.Walk(operand, func(dir string, info os.FileInfo, err error) error {
					if info.IsDir() && em.Empty(dir) {
						if err = rmdir(operand); err != nil {
							return err
						}
					}
					return nil
				}); err != nil {
					return err
				}
			}

			if em.Empty(operand) {
				if err = rmdir(operand); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
