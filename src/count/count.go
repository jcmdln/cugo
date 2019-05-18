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
	var (
		err   error
		index int
		total int
	)

	if len(operands) < 1 {
		err = errors.New("count: too few operands")
		return err
	}

	for index = range operands {
		total = index + 1
	}

	if _, err = fmt.Printf("%v\n", total); err != nil {
		return err
	}

	return nil
}
