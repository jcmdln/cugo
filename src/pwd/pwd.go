// SPDX-License-Identifier: ISC

package pwd

import (
	"fmt"
	"os"
	"path/filepath"
)

type Option struct {
	L bool
	P bool
}

func (opt *Option) Pwd() error {
	var (
		cwd string
		dir string
		err error
	)

	if !opt.L || opt.P {
		if cwd, err = os.Getwd(); err != nil {
			return err
		}

		if dir, err = filepath.EvalSymlinks(cwd); err != nil {
			return err
		}
	} else {
		dir = os.Getenv("PWD")
	}

	if _, err = fmt.Printf("%s\n", dir); err != nil {
		return err
	}

	return nil
}
