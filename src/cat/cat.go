package cat

import (
	"fmt"
	"io"
	"os"
)

type CatOptions struct {
	Unbuffered bool
}

type CatOpts func(*CatOptions)

func Unbuffered(unbuffered bool) CatOpts {
	return func(opts *CatOptions) {
		opts.Unbuffered = unbuffered
	}
}

func (c *CatOptions) Cat(operands []string) {
	var (
		operand string
		file    io.Reader
		buffer  []byte
		buff    = make([]byte, 1)
		err     error
	)

	if c.Unbuffered {
		buffer = buff
	} else {
		buffer = make([]byte, 4096)
	}

	if len(operands) == 0 {
		io.CopyBuffer(os.Stdout, os.Stdin, buffer)
	} else {
		for _, operand = range operands {
			if file, err = os.Open(operand); err != nil {
				fmt.Printf("cugo: %s", err)
				os.Exit(1)
			}

			if _, err = file.Read(buff); err != nil {
				fmt.Printf("cugo: %s", err)
				os.Exit(1)
			}

			io.CopyBuffer(os.Stdout, file, buffer)
		}
	}

	os.Exit(0)
}
