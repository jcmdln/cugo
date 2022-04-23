// SPDX-License-Identifier: ISC

package cat

import (
	"io"
	"os"
)

type Option struct {
	Unbuffered bool
}

func (opt *Option) Cat(files []*os.File) error {
	var err error

	for _, file := range files {
		if opt.Unbuffered {
			_, err = io.CopyBuffer(os.Stdout, file, make([]byte, 1))
		} else {
			_, err = io.Copy(os.Stdout, file)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
