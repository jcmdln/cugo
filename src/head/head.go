// head.go
//
// ISC License
//
// Copyright (c) 2022 Johnathan C. Maudlin <jcmdln@gmail.com>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE
// OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

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
