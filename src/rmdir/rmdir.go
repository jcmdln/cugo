// rmdir.go
//
// ISC License
//
// Copyright (c) 2022 Johnathan C. Maudlin <jcmdln@gmail.com>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE
// OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

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
