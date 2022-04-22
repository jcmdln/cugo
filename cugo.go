// cugo.go
//
// ISC License
//
// Copyright (c) 2022 Johnathan C. Maudlin <jcmdln@gmail.com>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE
// OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

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
	var err error

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

	c := cmd.Commands[command]
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
