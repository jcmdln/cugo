// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

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
		err     error
		file    io.Reader
		index   int
		line    string
		operand string
		read    *bufio.Reader
	)

	if opt.Number < 0 {
		err = errors.New("head: NUMBER can not be less than zero")
		return err
	}

	readlines := func(t *bufio.Reader) error {
		for index = 1; index <= opt.Number; index++ {
			if line, err = t.ReadString('\n'); err != nil && err != io.EOF {
				return err
			}

			if _, err = fmt.Printf(line); err != nil {
				return err
			}
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

	return nil
}
