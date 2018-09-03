package basename

import (
	"fmt"
	"path/filepath"
)

func Basename(args []string) {
	fmt.Printf("%s\n", filepath.Base(args[0]))
}
