// SPDX-License-Identifier: ISC

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

func usage() {
	var (
		err         error
		commands    []string
		currentLine string
		termWidth   int
	)

	if termWidth, _, err = term.Size(int(os.Stdin.Fd())); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("usage: cugo [command] [options]\n\n")
	fmt.Println("Subcommands:")
	for k := range cmd.Commands {
		commands = append(commands, k)
	}
	sort.Strings(commands)

	for index, command := range commands {
		if len(currentLine)+len(command) >= termWidth {
			fmt.Println(currentLine)
			currentLine = ""
		}

		currentLine += command
		if index != len(commands)-1 {
			currentLine += " "
		}
	}
	fmt.Printf("%s\n", currentLine)
}

func main() {
	var (
		c   cmd.Command
		err error
		ok  bool
	)

	if len(os.Args) < 2 {
		fmt.Println("error: missing subcommand")
		usage()
		os.Exit(1)
	}

	command := os.Args[1]
	if command == "-h" {
		usage()
		os.Exit(0)
	}

	if c, ok = cmd.Commands[command]; !ok {
		fmt.Println("error: no such command")
		usage()
		os.Exit(1)
	}

	f := c.Init()
	if err = f.Parse(os.Args[2:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = c.Run(f.Args()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
