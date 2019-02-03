// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package chmod - change file mode bits.
package chmod

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	ex "github.com/jcmdln/cugo/lib/exists"
)

var (
	// Recursive is a bool that when true changes files and directories
	// recursively.
	Recursive bool
)

// Chmod changes the file mode bits of provided files as specified by
// the mode operand. The mode of a file dictates its permissions, among
// other attributes. Currently the only supported mode operand uses
// octal numbers from 0 to 7.
func Chmod(args []string) {
	if len(args) < 2 {
		fmt.Printf("cugo: chmod: wrong number of arguments\n")
		os.Exit(1)
	}

	mode, err := strconv.ParseUint(args[0], 8, 32)
	if err != nil {
		fmt.Printf("cugo: %s\n", err)
		os.Exit(1)
	}

	for _, target := range args[1:] {
		if ex.Exists(target) {
			if Recursive {
				filepath.Walk(target, func(t string, info os.FileInfo, err error) error {
					err = os.Chmod(t, os.FileMode(mode))
					if err != nil {
						fmt.Printf("cugo: %s\n", err)
						os.Exit(1)
					}

					return nil
				})
			} else {
				err := os.Chmod(target, os.FileMode(mode))
				if err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				}
			}
		}
	}
}
