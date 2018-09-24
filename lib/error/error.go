package error

import (
	"fmt"
	"os"
)

func Error(prefix string, err error) bool {
	if err != nil {
		fmt.Printf("%s: %s\n", prefix, err)
		os.Exit(1)
	}

	return false
}
