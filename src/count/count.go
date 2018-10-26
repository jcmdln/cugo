// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// count the number of elements of an array
//
// count prints the number of arguments passed to it, and accepts no
// options.
package count

import (
	"fmt"
	"os"
)

// count prints the number of arguments passed to it.
func Count(args []string) {
	if len(args) < 1 {
		os.Exit(1)
	}

	var total int
	for i, _ := range args {
		total = i + 1
	}

	fmt.Printf("%v\n", total)
}
