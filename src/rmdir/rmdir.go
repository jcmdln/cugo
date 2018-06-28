package rmdir

import (
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
	} else {
		return false
	}
}

func Rmdir(args []string) {
	for _, target := range args {
		for empty(target) == false {
			filepath.Walk(target,
				func(t string, info os.FileInfo, err error) error {
					if info.IsDir() && empty(t) == true {
						os.Remove(t)
					}
					if !info.IsDir() {
						os.Remove(t)
					}
					return nil
				},
			)
		}

		if empty(target) == true {
			os.Remove(target)
		}
	}
}
