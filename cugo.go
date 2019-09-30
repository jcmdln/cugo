// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
)

func main() {
	if err := cmd.Command.Parse(os.Args[1:]); err != nil {
		var (
			s, c string
			l    []string
			i    int
		)

		l = cmd.Command.List()
		sort.Strings(l)

		for i, c = range l {
			s += c
			if i != len(l)-1 {
				s += " "
			}
		}

		fmt.Printf("%s\n\nAvailable commands:\n%s\n", err, s)
		os.Exit(1)
	}

	os.Exit(0)
}
