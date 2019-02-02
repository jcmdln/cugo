// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// return non-directory portion of a pathname
//
// basename receives the first argument and returns the last element of
// it's path. If the path is empty it returns ".", and if the path
// consists entirely of separators, basename returns a single separator.
//
// basename exits 0 on success, and >0 if an error occurs.
package basename

import (
	"fmt"
	"path/filepath"
)

func Basename(args []string) {
	fmt.Printf("%s\n", filepath.Base(args[0]))
}
