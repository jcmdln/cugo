package chmod

import (
	"fmt"
	"os"
	"strconv"

	er "github.com/jcmdln/cugo/lib/error"
	ex "github.com/jcmdln/cugo/lib/exists"
)

var (
	Changes   bool
	Quiet     bool
	Verbose   bool
	Recursive bool
)

func Chmod(args []string) {
	if len(args) < 2 {
		fmt.Printf("cugo: chmod: wrong number of arguments")
		os.Exit(1)
	}

	mode, err := strconv.ParseUint(args[0], 10, 32)
	er.Error("cugo", err)

	for _, target := range args[1:] {
		if ex.Exists(target) {
			err := os.Chmod(target, os.FileMode(mode))
			er.Error("cugo", err)
		}
	}
}
