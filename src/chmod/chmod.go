// chmod.go
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
