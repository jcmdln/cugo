package rmdir

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	Pathname bool
)

func empty(name string) bool {
	t, err := os.Open(name)
	defer t.Close()
	_, err = t.Readdirnames(1)
	if err == io.EOF {
		return true
	}

	return false
}

func Rmdir(args []string) {
	for _, target := range args {
		_, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Println("cugo: rm: Can't remove", "'"+target+"':",
				"no such file or directory")
			return
		}

		for !empty(target) {
			filepath.Walk(target,
				func(t string, info os.FileInfo, err error) error {
					if info.IsDir() && empty(t) {
						os.Remove(t)
					}
					if !info.IsDir() {
						os.Remove(t)
					}
					return nil
				},
			)
		}

		if !empty(target) {
			os.Remove(target)
		}
	}
}
