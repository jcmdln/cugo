// Copyright 2021 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that may
// be found in the LICENSE file.

package basename

import (
	"path/filepath"
	"strings"
)

// Basename returns the non-directory portion of a pathname by deleting
// any prefix ending with the last forward slash (‘/’), and a suffix if
// given which may be an empty string if not required.
func Basename(operand string, suffix string) (string, error) {
	var s string

	s = filepath.Base(operand)
	s = strings.Trim(s, suffix)

	return s, nil
}
