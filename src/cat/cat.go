// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cat

import (
	"io"
	"os"
)

func (opt *Options) Cat(operands []string) error {
	var (
		operand string
		file    io.Reader
		buffer  []byte
		buff    = make([]byte, 4096)
		unbuff  = make([]byte, 1)
		err     error
	)

	if opt.Unbuffered {
		buffer = unbuff
	} else {
		buffer = buff
	}

	if len(operands) == 0 {
		if _, err := io.CopyBuffer(os.Stdout, os.Stdin, buffer); err != nil {
			return err
		}
	} else {
		for _, operand = range operands {
			if file, err = os.Open(operand); err != nil {
				return err
			}

			if _, err = file.Read(buff); err != nil {
				return err
			}

			if _, err := io.CopyBuffer(os.Stdout, file, buffer); err != nil {
				return err
			}
		}
	}

	return nil
}
