// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// display first few lines of files
//
// SYNOPSIS
//
//     head [-n NUMBER] [FILE ...]
//
// DESCRIPTION
//
// head copies the first NUMBER of lines of each specified FILE to
// stdout. If no files are named, head copies lines from stdin. If
// NUMBER is omitted, it defaults to 10.
//
// The options are as follows:
//
//     -n        Copy the first NUMBER of lines of each file
//
// SEE ALSO
//
//     Not available.
//
// REFERENCES
//
//     http://man.openbsd.org/head
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/head.html
package head

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Options struct {
	Number int
}

type Opts func(*Options)

func Number(num int) Opts {
	return func(opt *Options) {
		opt.Number = num
	}
}

func (opt *Options) Head(operands []string) {
	var (
		f    io.Reader
		line string
		err  error
	)

	if opt.Number <= 0 {
		fmt.Println("cugo: head: ")
		os.Exit(1)
	}

	for _, operand := range operands {
		if f, err = os.Open(operand); err != nil {
			fmt.Println("cugo: head:", err)
			os.Exit(1)
		}

		read := bufio.NewReader(f)

		for i := 0; i <= opt.Number; i++ {
			if line, err = read.ReadString('\n'); err != nil {
				fmt.Println("cugo: head:", err)
				os.Exit(1)
			}

			fmt.Printf(line)
		}
	}

	os.Exit(0)
}
