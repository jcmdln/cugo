// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// remove directories
//
// rmdir removes the directory entry specified in each argument provided
// it is empty.
//
// Arguments are processed in the provided order and require both the
// parent directory and it's contained subdirectories to be empty.
//
// Available options:
//
//     -p, --parents    Remove parent directories.
//
//     -v, --verbose    Print a message when actions are taken.
//
package rmdir

import (
	"fmt"
	"os"
	"path/filepath"

	em "github.com/hlfstr/cugo/lib/empty"
	ex "github.com/hlfstr/cugo/lib/exists"
)

var (
	Parents bool
	Verbose bool
)

func rmdir(dir string) {
	err := os.Remove(dir)
	if err != nil {
		fmt.Printf("cugo: %s\n", err)
		os.Exit(1)
	} else {
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
