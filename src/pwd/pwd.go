// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pwd

import (
	"fmt"
	"os"
	"path/filepath"
)

func (opt *Options) Pwd() error {
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
