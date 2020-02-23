// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

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
