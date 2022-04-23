// SPDX-License-Identifier: ISC

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
