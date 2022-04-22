// chmod.go
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

package chmod

import (
	"os"
	"path/filepath"
	"strconv"

	ex "github.com/jcmdln/cugo/lib/exists"
)

func (opt *Options) Chmod(operands []string) error {
	var (
		err     error
		mode    os.FileMode
		modeVal uint64
		operand string
	)

	if modeVal, err = strconv.ParseUint(operands[0], 8, 32); err != nil {
		return err
	}

	mode = os.FileMode(modeVal)

	for _, operand = range operands[1:] {
		if err = ex.Exists(operand); err != nil {
			return err
		}

		if opt.Recursive {
			if err = filepath.Walk(operand, func(s string, i os.FileInfo, err error) error {
				if err = os.Chmod(s, mode); err != nil {
					return err
				}

				return nil
			}); err != nil {
				return err
			}
		} else {
			if err = os.Chmod(operand, mode); err != nil {
				return err
			}
		}
	}

	return nil
}
