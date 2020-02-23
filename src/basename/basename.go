// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package basename

import (
	"path/filepath"
	"strings"
)

// Basename returns the non-directory portion of a pathname by
// deleting any prefix ending with the last slash (‘/’), and a suffix
// if given which may be an empty string if not required.
func Basename(operand string, suffix string) (string, error) {
	var (
		s string
	)

	// If an empty string is passed as the operand, it causes a panic.
	// Catch whether the length of 'operand' is less than one, and
	// print a newline.
	if len(operand) < 1 {
		s = "\n"
	} else {
		s = filepath.Base(operand)

		if len(suffix) > 0 {
			s = strings.Trim(s, suffix)
		}

		s += "\n"
	}

	return s, nil
}
