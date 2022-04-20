// cugo.go
//
// Copyright 2021 Johnathan C. Maudlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

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
		command     string
		commands    []string
		currentLine string
		index       int
		termWidth   int
	)

	fmt.Printf("usage: cugo [command] [options]\n\n")
	fmt.Println("Subcommands:")
	for k := range cmd.Commands {
		commands = append(commands, k)
	}

	sort.Strings(commands)

	if termWidth, _, err = term.Size(int(os.Stdin.Fd())); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	for index, command = range commands {
		if len(currentLine)+len(command) >= termWidth {
			fmt.Printf("%s\n", currentLine)
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

	for k, v := range cmd.Commands {
		if command == k {
			f := v.Init()
			f.Parse(os.Args[2:])
			v.Run(f.Args())
		}
	}

	os.Exit(0)
}
