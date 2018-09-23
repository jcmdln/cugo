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

func touch(file string) {
	_, err := os.Create(file)
	if err != nil {
		fmt.Printf("cugo: %s\n", err)
	} else {
		if Verbose {
			fmt.Printf("cugo: touch: Created '%s'\n", file)
		}
	}
}

func Touch(args []string) {
	for _, file := range args {
		if !Create {
			touch(file)
		}
	}
}
