// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package head

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func (opt *Options) Head(operands []string) error {
	var (
		f    io.Reader
		nl   bool
		line string
		err  error
	)

	if opt.Number <= 0 {
		err = errors.New("head: NUMBER can not be less than zero")
		return err
	}

	for _, operand := range operands {
		if f, err = os.Open(operand); err != nil {
			return err
		}

		read := bufio.NewReader(f)

		for i := 0; i <= opt.Number; i++ {
			line, err = read.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					nl = true
				} else {
					fmt.Println("cugo: head:", err)
					return err
				}
			}

			fmt.Printf(line)
		}
	}

	if nl {
		fmt.Printf("\n")
	}

	return nil
}
