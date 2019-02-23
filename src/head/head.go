// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package head

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func (opt *Options) Head(operands []string) {
	var (
		f    io.Reader
		nl   bool
		line string
		err  error
	)

	if opt.Number <= 0 {
		fmt.Println("cugo: head: NUMBER can not be less than zero")
		os.Exit(1)
	}

	for _, operand := range operands {
		if f, err = os.Open(operand); err != nil {
			fmt.Println("cugo: head:", err)
			os.Exit(1)
		}

		read := bufio.NewReader(f)

		for i := 0; i <= opt.Number; i++ {
			line, err = read.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					nl = true
				} else {
					fmt.Println("cugo: head:", err)
					os.Exit(1)
				}
			}

			fmt.Printf(line)
		}
	}

	if nl {
		fmt.Printf("\n")
	}

	os.Exit(0)
}
