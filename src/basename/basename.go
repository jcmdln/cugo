package basename

import (
	"fmt"
	"path/filepath"
)

func Basename(args []string) {
	b := filepath.Base(args[0])
	fmt.Printf("%s\n", b)
}
