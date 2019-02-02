// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// create directories
//
// mkdir creates the directories named as arguments in the order
// specified using mode rwxrwxrwx (0777) as modified by the current
// umask.
//
// Available options:
//
//     -m, --mode       Set directory permissions to MODE value
//
//     -p, --parents    create missing parent directories
//
//     -v, --verbose    Display each directory after it was created
//
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
