// head.go
//
// Copyright 2021 Johnathan C. Maudlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

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
