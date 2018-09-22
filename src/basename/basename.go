package basename

import (
	"fmt"
	"path/filepath"
)

// basename takes any number of args and uses filepath.Base to return the
// last element of a path. To conform to POSIX behavior, it will only act
// on the first provided arg.
func Basename(args []string) {
	arg := args[0]
	fmt.Printf("%s\n", filepath.Base(arg))
}
