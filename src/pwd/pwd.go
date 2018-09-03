package pwd

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	L bool
	P bool
)

func Pwd() {
	var dir string

	if !L || P {
		d, _ := os.Getwd()
		dir, _ = filepath.EvalSymlinks(d)
	} else {
		dir = os.Getenv("PWD")
	}

	fmt.Printf("%s\n", dir)
}
