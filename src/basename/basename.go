// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package basename

import (
	"fmt"
	"os"
	"path/filepath"
)

func Basename(operand string) {
	fmt.Printf("%s\n", filepath.Base(operand))

	os.Exit(0)
}
