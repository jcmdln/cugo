// cat.go
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
