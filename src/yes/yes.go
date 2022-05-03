// SPDX-License-Identifier: ISC

package yes

import (
	"os"
	"strings"
)

func Yes(str string) {
	if len(str) == 0 {
		str = "y"
	}

	// Append newline character to string
	str = strings.Join([]string{str, "\n"}, "")

	for {
		os.Stdout.WriteString(str)
	}
}
