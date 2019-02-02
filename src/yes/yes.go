// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// repeatedly output specified string, or 'y'
//
// yes prints "y" unless a string to print is provided, forever.
//
package yes

import (
	"fmt"
	"strings"
)

func Yes(args []string) {
	if len(args) == 0 {
		for {
			fmt.Printf("y\n")
		}
	} else {
		out := strings.Join(args, " ")
		for {
			fmt.Printf("%s\n", out)
		}
	}
}
