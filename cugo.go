// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Unix utilities as a multi-call binary
//
// `cugo` provides common Unix core utilities in the form of a multi-call
// binary, with a focus on broad support for various Unix and Unix-like
// operating systems.
package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/jcmdln/cugo/cmd"
	"github.com/jcmdln/cugo/lib/term"
)

func main() {
	var width int

	if err := cmd.Command.Parse(os.Args[1:]); err != nil {
		commands := cmd.Command.List()
		sort.Strings(commands)

		fmt.Println("cugo:", err)
		fmt.Println("Available commands:")

		tw, _, _ := term.Size(int(os.Stdin.Fd()))
		for _, c := range commands {
			if len(c)+width > int(tw) {
				fmt.Printf("\n")
				width = 0
			}
			fmt.Printf("%s ", c)
			width += len(c) + len(" ")
		}

		fmt.Printf("\n")
		os.Exit(1)
	}

	os.Exit(0)
}
