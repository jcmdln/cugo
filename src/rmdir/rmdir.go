// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// remove directories
//
// SYNOPSIS
//
//     rmdir [-p] directory ...
//
// DESCRIPTION
//
// rmdir removes the directory entry specified by each directory
// argument, provided it is empty.
//
// The options are as follow:
//
//     -p        Remove empty parent directories.
//
// Arguments are processed in the provided order and will exit on any
// error, leaving following directories intact.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#Stat
//     https://golang.org/pkg/os/#Remove
//     https://golang.org/pkg/path/filepath/#Walk
//
// REFERENCES
//
//     http://man.openbsd.org/rmdir
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/rmdir.html
package rmdir

import (
	"fmt"
	"os"
	"path/filepath"

	em "github.com/jcmdln/cugo/lib/empty"
	ex "github.com/jcmdln/cugo/lib/exists"
)

type Options struct {
	Parents bool
	Verbose bool
}

type Opts func(*Options)

func Parents(parents bool) Opts {
	return func(opt *Options) {
		opt.Parents = parents
	}
}

func Verbose(verbose bool) Opts {
	return func(opt *Options) {
		opt.Verbose = verbose
	}
}

func (opt *Options) Rmdir(operands []string) {
	var (
		operand string
		err     error
	)

	rmdir := func(dir string) {
		if err = os.Remove(dir); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}

		if opt.Verbose {
			fmt.Printf("cugo: rm: Removed '%s'\n", dir)
		}
	}

	for _, operand = range operands {
		if !ex.Exists(operand) {
			fmt.Printf("cugo: rmdir: no such file or directory %s\n", operand)
			os.Exit(1)
		}

		if em.Empty(operand) {
			rmdir(operand)
		} else if opt.Parents {
			for !em.Empty(operand) {
				filepath.Walk(operand, func(dir string, info os.FileInfo, err error) error {
					if info.IsDir() && em.Empty(dir) {
						rmdir(dir)
					}

					return nil
				})
			}

			if em.Empty(operand) {
				rmdir(operand)
			}
		}
	}

	os.Exit(0)
}
