package ls

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	All       bool
	Recursive bool
)

func list(t string) {
	items, err := ioutil.ReadDir(t)
	if err != nil {
		fmt.Println("cugo: rm:", err)
	}

	for _, item := range items {
		if !All && strings.HasPrefix(item.Name(), ".") {
		} else {
			fmt.Printf(item.Name() + " ")
		}
	}

	fmt.Printf("\n")
}

func Ls(args []string) {
	if len(args) == 0 {
		list(".")
		return
	}

	for _, target := range args {
		_, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Printf("cugo: ls: '%s': No such file or directory\n", target)
			return
		}

		list(target)
	}
}
