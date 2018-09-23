package rmdir

import (
	"fmt"
	"os"
	"path/filepath"

	e "github.com/jcmdln/cugo/lib/empty"
	x "github.com/jcmdln/cugo/lib/exists"
)

var (
	Parents bool
	Verbose bool
)

func rmdir(dir string) {
	err := os.Remove(dir)
	if err != nil {
		fmt.Printf("cugo: %s\n", err)
	} else {
		if Verbose {
			fmt.Printf("cugo: rm: Removed '%s'\n", dir)
		}
	}
}

func Rmdir(args []string) {
	for _, dir := range args {
		if !x.Exists(dir) {
			fmt.Println("cugo: rmdir", dir+":",
				"no such file or directory")
			return
		}

		if e.Empty(dir) {
			rmdir(dir)
		} else if Parents {
			for !e.Empty(dir) {
				filepath.Walk(dir, func(t string, info os.FileInfo, err error) error {
					if info.IsDir() && e.Empty(t) {
						rmdir(t)
					}
					return nil
				})
			}
		}
	}
}
