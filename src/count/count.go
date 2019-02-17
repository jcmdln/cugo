// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package count

import (
	"fmt"
	"os"
)

// Count prints the number of arguments passed to it.
func Count(operands []string) {
	if len(operands) < 1 {
		os.Exit(1)
	}

	var total int
	for i := range operands {
		total = i + 1
	}

	fmt.Printf("%v\n", total)

	os.Exit(0)
}
