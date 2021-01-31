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

func main() {
	var (
		err         error
		command     string
		commandList []string
		currentLine string
		index       int
		termWidth   int
	)

	if err = cmd.Command.Parse(os.Args[1:]); err != nil {
		fmt.Printf("error: %s\n", err)

		if len(os.Args) < 2 {
			fmt.Printf("\nAvailable commands:\n")

			if termWidth, _, err = term.Size(int(os.Stdin.Fd())); err != nil {
				fmt.Print(err)
				os.Exit(1)
			}

			commandList = cmd.Command.List()
			sort.Strings(commandList)

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
		}

		os.Exit(1)
	}

	os.Exit(0)
}
