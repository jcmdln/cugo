// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mkdir

import (
	"fmt"
	"os"

	er "github.com/jcmdln/cugo/lib/error"
)

var (
	Mode    uint32
	Parents bool
	Verbose bool
)

func Mkdir(args []string) {
	for _, dir := range args {
		if Parents {
			err := os.MkdirAll(dir, os.FileMode(Mode))
			if !er.Error("cugo", err) {
				if Verbose {
					fmt.Printf("cugo: mkdir: Created %s\n", dir)
				}
			}
		} else {
			err := os.Mkdir(dir, os.FileMode(Mode))
			if !er.Error("cugo", err) {
				if Verbose {
					fmt.Printf("cugo: mkdir: Created %s\n", dir)
				}
			}
		}
	}
}
