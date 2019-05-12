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
		file    io.Reader
		read    *bufio.Reader
		newline bool
		line    string

		index int
		err   error
	)

	if opt.Number < 0 {
		err = errors.New("head: NUMBER can not be less than zero")
		return err
	}

	readlines := func(t *bufio.Reader) error {
		for index = 1; index <= opt.Number; index++ {
			line, err = read.ReadString('\n')

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
		read = bufio.NewReader(os.Stdout)
		if err := readlines(read); err != nil {
			return err
		}
	} else {
		for _, operand = range operands {
			if file, err = os.Open(operand); err != nil {
				return err
			}

			read = bufio.NewReader(file)
			if err := readlines(read); err != nil {
				return err
			}
		}
	}

	if newline {
		fmt.Printf("\n")
	}

	return nil
}
