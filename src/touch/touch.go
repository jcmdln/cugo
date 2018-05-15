package touch

import (
	"fmt"
	"os"
)

type TOUCH struct {
	// touchAccess   int
	// touchDate     string
	// touchModified int
	Create  bool
	Verbose bool
}

func Touch(args []string) {
	touch := &TOUCH{}

	Exists := func(t string) bool {
		_, err := os.Stat(t)
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	}

	Verbose := func(t string) {
		if touch.Verbose {
			fmt.Printf("cugo: touch: Created %s\n", t)
		}
	}

	for _, target := range args {
		if Exists(target) == false {
			if !touch.Create {
				os.Create(target)
				Verbose(target)
			}
		}
	}
}
