// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package rm

import (
	"fmt"
	"os"
	"path/filepath"

	em "github.com/jcmdln/cugo/lib/empty"
	pr "github.com/jcmdln/cugo/lib/prompt"
)

func (opt *Options) Rm(operands []string) error {
	var (
		err     error
		operand string
		ostat   os.FileInfo
	)

	remove := func(target string) error {
		if !opt.Force && opt.Interactive {
			if pr.Prompt("Remove '" + target + "'?") {
				if err = os.Remove(target); err != nil {
					err = fmt.Errorf("rm: %s: ", err)
					return err
				}
			}
		} else {
			if err = os.Remove(target); err != nil {
				err = fmt.Errorf("rm: %s: ", err)
				return err
			}
		}

		if opt.Verbose {
			fmt.Printf("cugo: rm: Removed '%s'\n", target)
		}

		return nil
	}

	for _, operand = range operands {
		if ostat, err = os.Stat(operand); os.IsNotExist(err) {
			err = fmt.Errorf("rm: %s: no such file or directory", operand)
			return err
		}

		if ostat.IsDir() {
			if opt.Dir && em.Empty(operand) {
				if err = remove(operand); err != nil {
					return err
				}
			}

			if opt.Recursive {
				for !em.Empty(operand) {
					if err = filepath.Walk(operand, func(target string, info os.FileInfo, err error) error {
						if info.IsDir() {
							if em.Empty(target) {
								if err = remove(operand); err != nil {
									return err
								}
							}
						} else {
							if err = remove(operand); err != nil {
								err = fmt.Errorf("rm: %s: directory not empty", err)
								return err
							}
						}

						return nil
					}); err != nil {
						//err = fmt.Errorf("rm: walking filepath failed")
						return err
					}
				}

				if em.Empty(operand) {
					if err = remove(operand); err != nil {
						err = fmt.Errorf("rm: %s", err)
						return err
					}
				}
			} else {
				err = fmt.Errorf("rm: %s: is a directory", operand)
				return err
			}
		} else {
			if err = remove(operand); err != nil {
				err = fmt.Errorf("rm: %s: ", err)
				return err
			}
		}
	}

	return nil
}
