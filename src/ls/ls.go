package ls

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type LS struct {
	All         bool
	Interactive bool
	Recursive   bool
	Verbose     bool
}

func Ls(args []string) {
	ls := &LS{}

	List := func(t string) {
		items, err := ioutil.ReadDir(t)
		if err != nil {
			fmt.Println("cugo: rm:", err)
		}
		for _, item := range items {
			if !ls.All && strings.HasPrefix(item.Name(), ".") {
			} else {
				fmt.Printf(item.Name() + " ")
			}
		}
		fmt.Printf("\n")
	}

	if len(args) == 0 {
		List(".")
		return
	}

	for _, target := range args {
		_, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Printf("cugo: ls: '%s': No such file or directory\n", target)
			return
		}
		List(target)
	}
}
