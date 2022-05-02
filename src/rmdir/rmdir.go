// SPDX-License-Identifier: ISC

package rmdir

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	em "github.com/jcmdln/cugo/lib/empty"
)

type Option struct {
	Parents bool
	Verbose bool
}

func (opt *Option) Rmdir(operands []string) error {
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
		if _, err = os.Stat(operand); err != nil {
			return errors.New("rmdir: " + operand + ": no such file or directory")
		}

		if em.Empty(operand) {
			if err = rmdir(operand); err != nil {
				return err
			}
		} else if opt.Parents {
			for !em.Empty(operand) {
				if err = filepath.Walk(operand, func(dir string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

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
