package rmdir

import (
	"fmt"
	"os"
	"path/filepath"

	e "github.com/jcmdln/cugo/lib/empty"
	x "github.com/jcmdln/cugo/lib/exists"
)

var (
	Pathname bool
	Verbose  bool
)

func rmdir(dir string) {
	os.Remove(dir)
	if Verbose {
		fmt.Printf("cugo: rm: Removed '%s'\n", dir)
	}

}

func Rmdir(args []string) {
	for _, dir := range args {
		if x.Exists(dir) {
			fmt.Println("cugo: rmdir", dir+":",
				"no such file or directory")
			return
		}

		if e.Empty(dir) {
			rmdir(dir)
			return
		}

		for !e.Empty(dir) {
			filepath.Walk(dir,
				func(t string, info os.FileInfo, err error) error {
					if info.IsDir() && e.Empty(t) {
						rmdir(t)
					}
					if !info.IsDir() {
						rmdir(t)
					}
					return nil
				},
			)
		}
	}
}
