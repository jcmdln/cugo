package touch

import (
	"fmt"
	"os"
)

var (
	// touchAccess   int
	// touchDate     string
	// touchModified int
	Create  bool
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
		fmt.Printf("cugo: touch: Created %s\n", t)
	}
}

func Touch(args []string) {
	for _, target := range args {
		if exists(target) == false {
			if !Create {
				os.Create(target)
				verbose(target)
			}
		}
	}
}
