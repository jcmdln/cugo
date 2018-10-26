// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// change file mode bits.
//
// This utility is going through some growing pains as Linux and Unix
// have heavily diverged on it's options. I'm leaning toward sticking
// with the defined behavior in OpenBSD's manual.
//
// Chmod changes the file mode bits of provided files as specified by
// the mode operand. The mode of a file dictates its permissions, among
// other attributes. Currently the only supported mode operand uses
// octal numbers from 0 to 7.
//
// Available options:
//
//     -h, --modify       treat symlinks like files; modify without
//                        following. -h and -R are mutually exclusive.
//
//     -H, --head         if -R is issued, follow symlinks but not
//                        symlinks found during traversal.
//
//     -L, --links        if -R is issued, all symlinks are followed.
//
//     -P, --physical     if -R is issued, no symlinks are followed.
//
//     -R, --recursive    change mode of the directory and it's contents.
//
package chmod

import (
	"fmt"
	"os"
	"strconv"

	er "github.com/jcmdln/cugo/lib/error"
	ex "github.com/jcmdln/cugo/lib/exists"
)

var (
	// Changes is a bool that when true reports when changes are made.
	Changes bool

	// Recursive is a bool that when true changes files and directories
	// recursively.
	Recursive bool

	// Quiet is a bool that when true suppresses most error messages.
	Quiet bool

	// Verbose is a bool that when true reports each processed file.
	Verbose bool
)

func Chmod(args []string) {
	if len(args) < 2 {
		fmt.Printf("cugo: chmod: wrong number of arguments")
		os.Exit(1)
	}

	mode, err := strconv.ParseUint(args[0], 8, 32)
	er.Error("cugo", err)

	for _, target := range args[1:] {
		if ex.Exists(target) {
			err := os.Chmod(target, os.FileMode(mode))
			er.Error("cugo", err)
		}
	}
}
