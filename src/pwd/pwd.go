// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// return working directory name
//
// SYNOPSIS
//
//     pwd [-LP]
//
// DESCRIPTION
//
// The pwd utility prints the absolute pathname of the current working
// directory to the standard output.
//
// The options are as follows:
//
//     -L        Print the logical (including symlinks) path of the
//               current directory.
//
//     -P        Print the physical path of the current directory.
//
// REFERENCES
//
//
package pwd

import (
	"fmt"
	"os"
	"path/filepath"
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
		if err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}

		dir, err = filepath.EvalSymlinks(cwd)
		if err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}
	} else {
		dir = os.Getenv("PWD")
	}

	fmt.Printf("%s\n", dir)
}
