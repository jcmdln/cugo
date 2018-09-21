package chmod

import (
	"fmt"
	"os"

	e "github.com/jcmdln/cugo/lib/exists"
)

var (
	Changes   bool
	Quiet     bool
	Verbose   bool
	Recursive bool
	Mode      uint32
)

func Chmod(args []string) {
	for _, target := range args {
		if e.Exists(target) {
			err := os.Chmod(target, os.FileMode(Mode))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
