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

type Options struct {
	Recursive bool
}

type Opts func(*Options)

func (opt *Options) Chmod(operands []string) {
	var (
		modeVal uint64
		mode    os.FileMode
		err     error
	)

	if modeVal, err = strconv.ParseUint(operands[0], 8, 32); err != nil {
		fmt.Printf("cugo: %s\n", err)
		os.Exit(1)
	}
	mode = os.FileMode(modeVal)

	for _, operand := range operands[1:] {
		if ex.Exists(operand) {
			if opt.Recursive {
				filepath.Walk(operand, func(t string, info os.FileInfo, err error) error {
					if err = os.Chmod(t, mode); err != nil {
						fmt.Printf("cugo: %s\n", err)
						os.Exit(1)
					}

					return nil
				})
			} else {
				if err = os.Chmod(operand, mode); err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				}
			}
		}
	}

	os.Exit(0)
}
