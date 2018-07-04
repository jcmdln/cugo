package mkdir

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	Mode    uint32
	Parents bool
	Verbose bool
)

func exists(t string) bool {
	_, err := os.Stat(t)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func verbose(t string) {
	if Verbose {
		fmt.Printf("cugo: mkdir: Created %s\n", t)
	}
}

func Mkdir(args []string) {
	for _, target := range args {
		if Parents {
			c := ""
			t := strings.Split(filepath.Clean(target), "/")

			for i := range t {
				c += t[i]
				if exists(c) == false {
					os.Mkdir(c, os.FileMode(Mode))
					verbose(c)
				}
				c += "/"
			}
		} else {
			if exists(filepath.Dir(target)) == true {
				os.Mkdir(target,
					os.FileMode(Mode))
				verbose(target)
			} else {
				fmt.Printf("cugo: mkdir: '%s' doesn't exist: "+
					"missing parent directory.\n", filepath.Dir(target))
			}
		}
	}
}
