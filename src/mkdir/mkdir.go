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

type Options struct {
	Mode    uint
	Parents bool
	Verbose bool
}

type Opts func(*Options)

func Mode(mode uint) Opts {
	return func(opt *Options) {
		opt.Mode = mode
	}
}

func Parents(unbuffered bool) Opts {
	return func(opt *Options) {
		opt.Parents = unbuffered
	}
}

func Verbose(verbose bool) Opts {
	return func(opt *Options) {
		opt.Verbose = verbose
	}
}

func (opt *Options) Mkdir(args []string) {
	var (
		dir  string
		mode os.FileMode
		err  error
	)

	mode = os.FileMode(uint32(opt.Mode))

	for _, dir = range args {
		if opt.Parents {
			if err = os.MkdirAll(dir, mode); err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			}

			if opt.Verbose {
				fmt.Printf("cugo: mkdir: Created %s\n", dir)
			}
		} else {
			if err = os.Mkdir(dir, mode); err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			}

			if opt.Verbose {
				fmt.Printf("cugo: mkdir: Created %s\n", dir)
			}
		}
	}

	os.Exit(0)
}
