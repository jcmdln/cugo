package touch

import (
	"fmt"
	"os"
)

var (
	Access    bool
	Create    bool
	Modified  bool
	Reference string
	Time      string
	Verbose   bool
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
		if !exists(target) {
			if !Create {
				os.Create(target)
				verbose(target)
			}
		}
	}
}
