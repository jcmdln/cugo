// cat.go
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
