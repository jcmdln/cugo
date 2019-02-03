// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// change file mode bits.
//
// SYNOPSIS
//
//     chmod [-R] MODE FILE ...
//
// DESCRIPTION
//
// Chmod changes the file mode bits of provided files as specified by
// the mode operand. The mode of a file dictates its permissions, among
// other attributes. Currently the only supported mode operand uses
// octal numbers from 0 to 7.
//
// The options are as follows:
//
//     -R        Change files and directories recursively.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#Chmod
//
// REFERENCES
//
// 	   http://man.openbsd.org/chmod
// 	   http://pubs.opengroup.org/onlinepubs/9699919799/utilities/chmod.html
package chmod

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	ex "github.com/jcmdln/cugo/lib/exists"
)

var (
	// Recursive is a boolean that, when true, specifies to recursively
	// change the mode of the directory and all it's children.
	Recursive bool
)

// Chmod receives an array of strings, the first of which must be the
// MODE operand, followed by the files to have their permission bits set
// to MODE. If the first argument is not an octal representation of a
// MODE, Chmod will exit with a non-zero exit code.
func Chmod(args []string) {
	if len(args) < 2 {
		fmt.Printf("cugo: chmod: wrong number of arguments\n")
		os.Exit(1)
	}

	m, err := strconv.ParseUint(args[0], 8, 32)
	if err != nil {
		fmt.Printf("cugo: %s\n", err)
		os.Exit(1)
	}

	mode := os.FileMode(m)

	for _, target := range args[1:] {
		if ex.Exists(target) {
			if Recursive {
				filepath.Walk(target, func(t string, info os.FileInfo, err error) error {
					err = os.Chmod(t, mode)
					if err != nil {
						fmt.Printf("cugo: %s\n", err)
						os.Exit(1)
					}

					return nil
				})
			} else {
				err := os.Chmod(target, mode)
				if err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				}
			}
		}
	}
}
