// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package count

import "fmt"

func Count(args []string) {
	var total int
	for i, _ := range args {
		total = i + 1
	}
	fmt.Printf("%v\n", total)
}
