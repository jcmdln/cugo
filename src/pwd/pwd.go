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
// SEE ALSO
//
//     https://golang.org/pkg/os/#Getwd
//
// REFERENCES
//
//     http://man.openbsd.org/pwd
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pwd.html
package pwd

import (
	"fmt"
	"os"
	"path/filepath"
)

type Options struct {
	L bool
	P bool
}

type Opts func(*Options)

func L(l bool) Opts {
	return func(opt *Options) {
		opt.L = l
	}
}

func P(p bool) Opts {
	return func(opt *Options) {
		opt.P = p
	}
}

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
