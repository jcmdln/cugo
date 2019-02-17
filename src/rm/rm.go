// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// remove files and directories
//
// SYNOPSIS
//
//     rm [-dfiRrv] FILE ...
//
// DESCRIPTION
//
// rm attempts to remove non-directory type files specified as arguments
// regardless of the file's permissions.
//
// The options are as follows:
//
//     -d        Remove empty directories.
//
//     -f        Skip prompts and ignore warnings.
//
//     -i        Prompt before each removal.
//
//     -R, -r    Remove directories and their contents recursively.
//
//     -v        Print a message when actions are taken.
//
// It is an error to attempt to remove the root directory or the files
// “.” or “..”. It is forbidden to remove the file “..” merely to avoid
// the antisocial consequences of inadvertently doing something like
// “rm -r .*”.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#Stat
//     https://golang.org/pkg/os/#Remove
//     https://golang.org/pkg/path/filepath/#Walk
//
// REFERENCES
//
//     http://man.openbsd.org/rm
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/rm.html
package rm

import (
	"fmt"
	"os"
	"path/filepath"

	em "github.com/jcmdln/cugo/lib/empty"
	pr "github.com/jcmdln/cugo/lib/prompt"
)

type Options struct {
	Dir         bool
	Force       bool
	Interactive bool
	Recursive   bool
	Verbose     bool
}

type Opts func(*Options)

func Dir(dir bool) Opts {
	return func(opt *Options) {
		opt.Dir = dir
	}
}

func Force(force bool) Opts {
	return func(opt *Options) {
		opt.Force = force
	}
}

func Interactive(interactive bool) Opts {
	return func(opt *Options) {
		opt.Interactive = interactive
	}
}

func Recursive(recursive bool) Opts {
	return func(opt *Options) {
		opt.Recursive = recursive
	}
}

func Verbose(verbose bool) Opts {
	return func(opt *Options) {
		opt.Verbose = verbose
	}
}

func (opt *Options) Rm(operands []string) {
	var (
		operand string
		ostat   os.FileInfo
		err     error
	)

	remove := func(target string) {
		if !opt.Force && opt.Interactive {
			if pr.Prompt("Remove '" + target + "'?") {
				if err = os.Remove(target); err != nil {
					fmt.Printf("cugo: rm: %s\n", err)
					os.Exit(1)
				}

				if opt.Verbose {
					fmt.Printf("cugo: rm: Removed '%s'\n", target)
				}
			}
		} else {
			if err = os.Remove(target); err != nil {
				fmt.Printf("cugo: rm: %s\n", err)
				os.Exit(1)
			}

			if opt.Verbose {
				fmt.Printf("cugo: rm: Removed '%s'\n", target)
			}
		}
	}

	for _, operand = range operands {
		if ostat, err = os.Stat(operand); os.IsNotExist(err) {
			fmt.Printf("cugo: rm %s: no such file or directory\n", operand)
			os.Exit(1)
		}

		if ostat.IsDir() {
			if opt.Dir && em.Empty(operand) {
				remove(operand)
			}

			if opt.Recursive {
				for !em.Empty(operand) {
					filepath.Walk(operand, func(target string, info os.FileInfo, err error) error {
						if info.IsDir() {
							if em.Empty(target) {
								remove(target)
							}
						} else {
							remove(target)
						}

						return nil
					})
				}

				if em.Empty(operand) {
					remove(operand)
				}
			} else {
				fmt.Printf("cugo: rm: '%s' is a directory\n", operand)
			}
		} else {
			remove(operand)
		}
	}

	os.Exit(0)
}
