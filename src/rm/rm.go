// rm.go
//
// Copyright 2021 Johnathan C. Maudlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

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
