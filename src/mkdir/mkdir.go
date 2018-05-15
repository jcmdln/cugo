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

func Mkdir(args []string) {
	Exists := func(t string) bool {
		_, err := os.Stat(t)
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	}

	Verbose := func(t string) {
		if Verbose {
			fmt.Printf("cugo: mkdir: Created %s\n", t)
		}
	}

	for _, target := range args {
		if Parents {
			c := "."
			t := strings.Split(filepath.Clean(target), "/")
			for i := range t {
				c += "/" + t[i]
				if Exists(c) == false {
					os.Mkdir(c, os.FileMode(Mode))
					Verbose(c)
				}
			}
		} else {
			if Exists(filepath.Dir(target)) == true {
				os.Mkdir(target,
					os.FileMode(Mode))
				Verbose(target)
			} else {
				fmt.Printf("cugo: mkdir: '%s' doesn't exist: "+
					"missing parent directory.\n", filepath.Dir(target))
			}
		}
	}
}
