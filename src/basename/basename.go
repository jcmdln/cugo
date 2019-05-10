// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package basename

import (
	"fmt"
	"path/filepath"
)

// Basename receives any number of operands, but only passes the first
// item of the []string to filepath.Base to mimic the behavior of other
// implementations.

func Basename(operand []string) error {
	// If an empty string is passed, it causes a panic. Catch whether
	// the length of 'operand' is 0, and print a newline.
	if len(operand) < 1 {
		fmt.Println("")
	} else {
		s := filepath.Base(operand[0])
		if _, err := fmt.Printf("%s\n", s); err != nil {
			return err
		}
	}

	return nil
}
