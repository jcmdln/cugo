// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rmdir

import (
	"fmt"
	"os"
	"path/filepath"

	em "github.com/jcmdln/cugo/lib/empty"
	er "github.com/jcmdln/cugo/lib/error"
	ex "github.com/jcmdln/cugo/lib/exists"
)

var (
	Parents bool
	Verbose bool
)

func rmdir(dir string) {
	err := os.Remove(dir)
	if !er.Error("cugo", err) {
		if Verbose {
			fmt.Printf("cugo: rm: Removed '%s'\n", dir)
		}
	}
}

func Rmdir(args []string) {
	for _, dir := range args {
		if !ex.Exists(dir) {
			fmt.Println("cugo: rmdir", dir+":",
				"no such file or directory")
			return
		}

		if em.Empty(dir) {
			rmdir(dir)
		} else if Parents {
			for !em.Empty(dir) {
				filepath.Walk(dir, func(d string, info os.FileInfo, err error) error {
					if info.IsDir() && em.Empty(d) {
						rmdir(d)
					}

					return nil
				})
			}

			if em.Empty(dir) {
				rmdir(dir)
			}
		}
	}
}
