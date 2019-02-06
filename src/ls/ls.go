// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// list directory contents
//
// SYNOPSIS
//
//     ls [-aR] [FILE ...]
//
// DESCRIPTION
//
// ls receives operands and may display the name (or other information)
// of the provided file or directory. For each named directory, ls
// displays the names of files contained within that directory as well
// as any requested, associated information.
//
// The options are as follows:
//
//     -a        Include directory entries that start with a dot ('.').
//
//     -R        Recursively list encountered subdirectories.
//
// SEE ALSO
//
//     https://golang.org/pkg/io/ioutil/#ReadDir
//
// REFERENCES
//
//     http://man.openbsd.org/ls
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/ls.html
package ls

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	ex "github.com/jcmdln/cugo/lib/exists"
)

var (
	All       bool
	Recursive bool

	operand string
	items   []os.FileInfo
	item    os.FileInfo
	err     error
)

func list(target string) {
	if items, err = ioutil.ReadDir(target); err != nil {
		fmt.Println("cugo: rm:", err)
		os.Exit(1)
	}

	for _, item = range items {
		if !All && strings.HasPrefix(item.Name(), ".") {
		} else {
			fmt.Printf(item.Name() + " ")
		}
	}

	fmt.Printf("\n")
}

// Ls will list all targets and their children in the order provided,
// and will recursively list children if specified.
func Ls(args []string) {
	if len(args) == 0 {
		list(".")
	} else {
		for _, operand = range args {
			if !ex.Exists(operand) {
				fmt.Printf("cugo: ls: '%s': No such file or directory\n", operand)
				os.Exit(1)
			}

			list(operand)
		}
	}

	os.Exit(0)
}
