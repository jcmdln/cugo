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
		commands    []string
		currentLine string
	)

	termWidth, _, err := term.Size(int(os.Stdin.Fd()))
	if err != nil {
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

	fmt.Println(currentLine)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: missing subcommand")
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-h":
		usage()
		os.Exit(0)
	}

	command := cmd.Commands[os.Args[1]]
	if command == nil {
		fmt.Printf("error: no such command '%s'\n", os.Args[1])
		usage()
		os.Exit(1)
	}

	flags := command.Init()
	if err := flags.Parse(os.Args[2:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := command.Run(flags.Args()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
