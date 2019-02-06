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

	ex "github.com/jcmdln/cugo/lib/exists"
)

var (
	// Recursive is a boolean that, when true, specifies to recursively
	// change the mode of the directory and all it's children.
	Recursive bool

	mode os.FileMode
	err  error
)

// Chmod receives a MODE, the MODE operand, and the files to have their
// permission bits set to MODE.
func Chmod(mode os.FileMode, files []string) {
	for _, target := range files {
		if ex.Exists(target) {
			if Recursive {
				filepath.Walk(target, func(t string, info os.FileInfo, err error) error {
					if err = os.Chmod(t, mode); err != nil {
						fmt.Printf("cugo: %s\n", err)
						os.Exit(1)
					}

					return nil
				})
			} else {
				if err = os.Chmod(target, mode); err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				}
			}
		}
	}

	os.Exit(0)
}
