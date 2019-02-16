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
	"strings"

	"github.com/jcmdln/cugo/cmd"
)

func main() {
	if err := cmd.Command.Parse(os.Args[1:]); err != nil {
		commands := cmd.Command.List()
		sort.Strings(commands)

		fmt.Printf("cugo: %s\nAvailable commands: %s\n",
			err, strings.Join(commands, " "))
	}
}
