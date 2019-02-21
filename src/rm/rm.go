// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rm

import (
	"fmt"
	"os"
	"path/filepath"

	em "github.com/jcmdln/cugo/lib/empty"
	pr "github.com/jcmdln/cugo/lib/prompt"
)

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
