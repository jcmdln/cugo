// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// repeatedly affirmative
//
// SYNOPSIS
//
//     yes STRING ...
//
// DESCRIPTION
//
// yes prints "y" or a provided string. Forever. If a series of strings
// has been provided, they will be treated as one string.
//
// SEE ALSO
//
//     Nothing to display.
//
// REFERENCES
//
//     http://man.openbsd.org/yes
//     https://linux.die.net/man/1/yes
package yes

import (
	"fmt"
	"strings"
)

func Yes(operands []string) {
	if len(operands) == 0 {
		for {
			fmt.Printf("y\n")
		}
	} else {
		out := strings.Join(operands, " ")
		for {
			fmt.Printf("%s\n", out)
		}
	}
}
