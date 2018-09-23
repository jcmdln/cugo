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

// if Force {
// 	os.Remove(t)
// 	if Verbose {
// 		fmt.Printf("cugo: rm: Removed '%s'\n", t)
// 	}
// }

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
		t, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Printf("cugo: rm %s: no such file or directory\n", target)
			return
		}

		if t.IsDir() {
			if Dir && e.Empty(target) {
				rm(target)
				return
			}

			if Recursive {
				for !e.Empty(target) {
					filepath.Walk(target, func(t string, info os.FileInfo, err error) error {
						if info.IsDir() && e.Empty(t) {
							rm(t)
						}
						if !info.IsDir() {
							rm(t)
						}
						return nil
					})
				}

				if e.Empty(target) {
					rm(target)
				}
			} else {
				fmt.Printf("cugo: rm %s: is a directory", target)
			}
		} else {
			rm(target)
		}
	}
}
