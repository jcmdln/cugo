package rm

import (
	"fmt"
	"os"
	"path/filepath"

	e "github.com/jcmdln/cugo/lib/empty"
	p "github.com/jcmdln/cugo/lib/prompt"
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
		fmt.Printf("cugo: %s\n", err)
	}
	if Verbose {
		fmt.Printf("cugo: rm: Removed '%s'\n", target)
	}
}

func remove(target string) {
	if !Force && Interactive {
		if p.Prompt("Remove '" + target + "'?") {
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
			if Dir && e.Empty(target) {
				remove(target)
				return
			}

			if Recursive {
				for !e.Empty(target) {
					filepath.Walk(target, func(t string, info os.FileInfo, err error) error {
						if info.IsDir() {
							if e.Empty(t) {
								remove(t)
							}
						} else {
							remove(t)
						}

						return nil
					})
				}

				if e.Empty(target) {
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
