// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package count

import (
	"errors"
	"fmt"
)

// Count prints the number of arguments passed to it.
func Count(operands []string) error {
	if len(operands) < 1 {
		err := errors.New("count: too few operands")
		return err
	}

	var total int
	for i := range operands {
		total = i + 1
	}
	fmt.Printf("%v\n", total)

	return nil
}
