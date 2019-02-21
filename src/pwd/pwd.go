// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pwd

import (
	"fmt"
	"os"
	"path/filepath"
)

func (opt *Options) Pwd() {
	var (
		cwd string
		dir string
		err error
	)

	if !opt.L || opt.P {
		if cwd, err = os.Getwd(); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}

		if dir, err = filepath.EvalSymlinks(cwd); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}
	} else {
		dir = os.Getenv("PWD")
	}

	fmt.Printf("%s\n", dir)

	os.Exit(0)
}
