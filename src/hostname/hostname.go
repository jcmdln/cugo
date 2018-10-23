// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hostname

import (
	"fmt"
	"os"

	er "github.com/jcmdln/cugo/lib/error"
)

func Hostname(args []string) {
	h, err := os.Hostname()
	if !er.Error("cugo", err) {
		fmt.Printf("%s\n", h)
	}
}
