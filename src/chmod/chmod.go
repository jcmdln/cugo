// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

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
