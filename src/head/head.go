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
		operand string
		newline bool
		line    string

		file io.Reader
		read *bufio.Reader

		index int
		err   error
	)

	if opt.Number < 0 {
		err = errors.New("head: NUMBER can not be less than zero")
		return err
	}

	readlines := func(t *bufio.Reader) error {
		for index = 1; index <= opt.Number; index++ {
			line, err = t.ReadString('\n')

			if err != nil {
				if err == io.EOF {
					newline = true
				} else {
					fmt.Println("cugo: head:", err)
					return err
				}
			}

			fmt.Printf(line)
		}

		return nil
	}

	if len(operands) < 1 {
		if _, err = os.Stdin.Stat(); err != nil {
			read = bufio.NewReader(os.Stdout)
		} else {
			read = bufio.NewReader(os.Stdin)
		}

		if err = readlines(read); err != nil {
			return err
		}
	} else {
		for _, operand = range operands {
			if file, err = os.Open(operand); err != nil {
				return err
			}

			read = bufio.NewReader(file)
			if err = readlines(read); err != nil {
				return err
			}
		}
	}

	if newline {
		fmt.Printf("\n")
	}

	return nil
}
