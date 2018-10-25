// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// return non-directory portion of a pathname
package basename

import (
	"fmt"
	"path/filepath"
)

// basename receives the first argument and returns the last element of
// it's path. If the path is empty it returns ".". If the path consists
// entirely of separators, basename returns a single separator.
func Basename(args []string) {
	arg := args[0]
	fmt.Printf("%s\n", filepath.Base(arg))
}
