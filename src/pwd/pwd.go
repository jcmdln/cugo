// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// return working directory name
//
// pwd prints the absolute pathname of the current working directory.
//
// Available options
//
//     -L, --links       Read current directory from the environment
//                       including symlinks.
//
//     -P, --physical    Absolute path of current directory without
//                       symlinks.
//
package pwd

import (
	"fmt"
	"os"
	"path/filepath"

	er "github.com/jcmdln/cugo/lib/error"
)

var (
	L bool
	P bool
)

func Pwd() {
	var (
		cwd string
		dir string
		err error
	)

	if !L || P {
		cwd, err = os.Getwd()
		er.Error("cugo", err)

		dir, err = filepath.EvalSymlinks(cwd)
		er.Error("cugo", err)
	} else {
		dir = os.Getenv("PWD")
	}

	fmt.Printf("%s\n", dir)
}
