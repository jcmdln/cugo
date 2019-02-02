// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// remove files and directories
//
// rm attempts to remove non-directory type files specified as arguments
// regardless of the file's permissions.
//
// Available options:
//
//     -d, --dir            Remove empty directories.
//
//     -f, --force          Skip prompts and ignore warnings.
//
//     -i, --interactive    Prompt before each removal.
//
//     -r, --recursive      Remove directories and their contents
//                          recursively.
//
//     -v, --verbose        Print a message when actions are taken.
//
package rm

import (
	"fmt"
	"os"
	"path/filepath"

	em "github.com/jcmdln/cugo/lib/empty"
	pr "github.com/jcmdln/cugo/lib/prompt"
)

var (
	Dir         bool
	Force       bool
	Interactive bool
	Recursive   bool
	Verbose     bool
)

func rm(target string) {
	err := os.Remove(target)
	if err != nil {
		fmt.Printf("cugo: rm: Removed '%s'\n", target)
		os.Exit(1)
	} else {
		if Verbose {
			fmt.Printf("cugo: rm: Removed '%s'\n", target)
		}
	}
}

func remove(target string) {
	if !Force && Interactive {
		if pr.Prompt("Remove '" + target + "'?") {
			rm(target)
		}
	} else {
		rm(target)
	}
}

func Rm(args []string) {
	for _, target := range args {
		cur, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Printf("cugo: rm %s: no such file or directory\n", target)
			return
		}

		if cur.IsDir() {
			if Dir && em.Empty(target) {
				remove(target)
				return
			}

			if Recursive {
				for !em.Empty(target) {
					filepath.Walk(target, func(t string, info os.FileInfo, err error) error {
						if info.IsDir() {
							if em.Empty(t) {
								remove(t)
							}
						} else {
							remove(t)
						}

						return nil
					})
				}

				if em.Empty(target) {
					remove(target)
				}
			} else {
				fmt.Printf("cugo: rm: '%s' is a directory\n", target)
			}
		} else {
			remove(target)
		}
	}
}
