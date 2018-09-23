package basename

import (
	"fmt"
	"path/filepath"
)

func Basename(args []string) {
	arg := args[0]
	fmt.Printf("%s\n", filepath.Base(arg))
}
