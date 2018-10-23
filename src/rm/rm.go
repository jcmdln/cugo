// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rm

import (
	"fmt"
	"os"
	"path/filepath"

	em "github.com/jcmdln/cugo/lib/empty"
	er "github.com/jcmdln/cugo/lib/error"
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
	if !er.Error("cugo", err) {
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
				fmt.Printf("cugo: rm %s: is a directory", target)
			}
		} else {
			remove(target)
		}
	}
}
