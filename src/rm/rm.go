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

func remove(t string) {
	if Force {
		os.Remove(t)
		if Verbose {
			fmt.Printf("cugo: rm: Removed '%s'\n", t)
		}
	} else if Interactive {
		if p.Prompt("Remove '" + t + "'?") {
			os.Remove(t)
			if Verbose {
				fmt.Printf("cugo: rm: Removed '%s'\n", t)
			}
		}
	}
}

func Rm(args []string) {
	for _, target := range args {
		t, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Printf("cugo: rm %s: no such file or directory", target)
			return
		}

		if t.IsDir() {
			if Dir && e.Empty(target) {
				remove(target)
				return
			}

			if Recursive {
				for !e.Empty(target) {
					filepath.Walk(target, func(t string, info os.FileInfo, err error) error {
						if info.IsDir() && e.Empty(t) {
							remove(t)
						}
						if !info.IsDir() {
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
