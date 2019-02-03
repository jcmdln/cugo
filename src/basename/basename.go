// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package basename - return non-directory portion of a pathname
package basename

import (
	"fmt"
	"path/filepath"
)

func Basename(args []string) {
	fmt.Printf("%s\n", filepath.Base(args[0]))
}
