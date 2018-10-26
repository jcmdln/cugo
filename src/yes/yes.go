// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// repeatedly output specified string, or 'y'
//
// yes prints "y" unless a string to print is provided, forever.
//
package yes

import (
	"io"
	"os"
	"strings"
)

func Yes(args []string) {
	if len(args) == 0 {
		for true {
			io.WriteString(os.Stdout, "y\n")
		}
	} else {
		for true {
			out := strings.Join(args, " ") + "\n"
			io.WriteString(os.Stdout, out)
		}
	}
}
