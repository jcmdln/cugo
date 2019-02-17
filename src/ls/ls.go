// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ls

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	ex "github.com/jcmdln/cugo/lib/exists"
	"github.com/jcmdln/cugo/lib/term"
)

type Options struct {
	All       bool
	Recursive bool
}

type Opts func(*Options)

// Ls will list all targets and their children in the order provided,
// and will recursively list children if specified.
func (opt *Options) Ls(operands []string) {
	var (
		operand string
		items   []os.FileInfo
		item    os.FileInfo
		err     error

		width int
	)

	list := func(target string) {
		if items, err = ioutil.ReadDir(target); err != nil {
			fmt.Println("cugo: rm:", err)
			os.Exit(1)
		}

		tw, _, _ := term.Size(int(os.Stdin.Fd()))
		for _, item = range items {
			if !opt.All && strings.HasPrefix(item.Name(), ".") {
			} else {
				if len(item.Name())+width > int(tw) {
					fmt.Printf("\n")
					width = 0
				}

				fmt.Printf("%s ", item.Name())
				width += len(item.Name()) + len(" ")
			}
		}

		fmt.Printf("\n")
	}

	if len(operands) == 0 {
		list(".")
	} else {
		for _, operand = range operands {
			if !ex.Exists(operand) {
				fmt.Printf("cugo: ls: '%s': No such file or directory\n", operand)
				os.Exit(1)
			}

			list(operand)
		}
	}

	os.Exit(0)
}
