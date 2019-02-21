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
