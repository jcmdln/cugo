// SPDX-License-Identifier: ISC

package chmod

import (
	"os"
	"path/filepath"
	"strconv"
)

type Option struct {
	Recursive bool
}

func (opt *Option) Chmod(mode string, files []*os.File) error {
	var (
		err     error
		modeVal uint64
	)

	if modeVal, err = strconv.ParseUint(mode, 8, 32); err != nil {
		return err
	}

	for _, file := range files {
		if _, err = file.Stat(); err != nil {
			return err
		}

		if opt.Recursive {
			if err = filepath.Walk(file.Name(), func(s string, i os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if err = os.Chmod(s, os.FileMode(modeVal)); err != nil {
					return err
				}

				return nil
			}); err != nil {
				return err
			}
		} else {
			if err = os.Chmod(file.Name(), os.FileMode(modeVal)); err != nil {
				return err
			}
		}
	}

	return nil
}
