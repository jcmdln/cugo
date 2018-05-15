package pwd

import (
	"fmt"
	"os"
	"path/filepath"
)

type PWD struct {
	L bool
	P bool
}

func Pwd() {
	pwd := &PWD{}

	var dir string

	if !pwd.P && pwd.L {
		dir = os.Getenv("PWD")
	} else if pwd.P {
		d, _ := os.Getwd()
		dir, _ = filepath.EvalSymlinks(d)
	} else {
		dir = os.Getenv("PWD")
	}

	fmt.Printf("%s\n", dir)
}
