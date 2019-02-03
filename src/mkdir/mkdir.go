// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// create directories
//
// SYNOPSIS
//
//     mkdir [-p] [-m MODE] DIRECTORY ...
//
// DESCRIPTION
//
// mkdir creates the directories named as arguments in the order
// specified using mode rwxrwxrwx (0777) as modified by the current
// umask.
//
// The options are as follows:
//
//     -m        Set the file permission bits of a new directory to MODE.
//
//     -p        Create missing parent directories as required.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#FileMode
//     https://golang.org/pkg/os/#Mkdir
//     https://golang.org/pkg/os/#MkdirAll
//
// REFERENCES
//
//     http://man.openbsd.org/mkdir
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/mkdir.html
package mkdir

import (
	"fmt"
	"os"
)

var (
	Mode    uint
	Parents bool
	Verbose bool
)

func Mkdir(args []string) {
	mode := os.FileMode(uint32(Mode))

	for _, dir := range args {
		if Parents {
			err := os.MkdirAll(dir, mode)
			if err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			} else {
				if Verbose {
					fmt.Printf("cugo: mkdir: Created %s\n", dir)
				}
			}
		} else {
			err := os.Mkdir(dir, mode)
			if err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			} else {
				if Verbose {
					fmt.Printf("cugo: mkdir: Created %s\n", dir)
				}
			}
		}
	}
}
