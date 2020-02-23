// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

// Unix utilities as a multi-call binary
//
// `cugo` provides common Unix core utilities in the form of a
// multi-call binary, with a focus on broad support for various Unix
// and Unix-like operating systems.
package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/jcmdln/cugo/cmd"
	"github.com/jcmdln/cugo/lib/term"
)

func main() {
	var (
		command     string
		commandList []string
		currentLine string
		index       int
		termWidth   int

		err error
	)

	if err = cmd.Command.Parse(os.Args[1:]); err != nil {
		fmt.Printf("error: %s\n", err)

		if termWidth, _, err = term.Size(int(os.Stdin.Fd())); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		commandList = cmd.Command.List()
		sort.Strings(commandList)

		fmt.Printf("\nAvailable commands:\n")

		for index, command = range commandList {
			if len(currentLine)+len(command) >= termWidth {
				fmt.Printf("%s\n", currentLine)
				currentLine = ""
			}

			currentLine += command
			if index != len(commandList)-1 {
				currentLine += " "
			}
		}

		fmt.Printf("%s\n", currentLine)

		os.Exit(1)
	}

	os.Exit(0)
}
