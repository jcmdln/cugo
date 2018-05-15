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

func Touch(args []string) {
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
			fmt.Printf("cugo: touch: Created %s\n", t)
		}
	}

	for _, target := range args {
		if Exists(target) == false {
			if !Create {
				os.Create(target)
				Verbose(target)
			}
		}
	}
}
