// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cat

import (
	"errors"
	"io"
	"os"
)

func (opt *Options) Cat(operands []string) error {
	var (
		buffer  []byte
		err     error
		file    io.Reader
		operand string
	)

	if opt.Unbuffered {
		buffer = make([]byte, 1)
	} else {
		buffer = make([]byte, 4096)
	}

	cat := func(in io.Reader) error {
		if _, err = io.CopyBuffer(os.Stdout, in, buffer); err != nil {
			err = errors.New("cat: /: is a directory")
			return err
		}

		return nil
	}

	if len(operands) == 0 {
		if err = cat(os.Stdin); err != nil {
			return err
		}
	} else {
		for _, operand = range operands {
			if file, err = os.Open(operand); err != nil {
				err = errors.New("cat: no such file or directory")
				return err
			}

			if err = cat(file); err != nil {
				return err
			}
		}
	}

	return nil
}
