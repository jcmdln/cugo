// rmdir.go
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
