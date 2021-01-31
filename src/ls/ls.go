// ls.go
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

package ls

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	ex "github.com/jcmdln/cugo/lib/exists"
	"github.com/jcmdln/cugo/lib/term"
)

// Ls will list all targets and their children in the order provided,
// and will recursively list children if specified.
func (opt *Options) Ls(operands []string) error {
	var (
		operand string
		index   int
		items   []os.FileInfo
		item    os.FileInfo
		entries string

		width int
		err   error
	)

	list := func(target string) error {
		if items, err = ioutil.ReadDir(target); err != nil {
			return err
		}

		tw, _, _ := term.Size(int(os.Stdout.Fd()))
		for index, item = range items {
			if !opt.All && strings.HasPrefix(item.Name(), ".") {
				// Do nothing
			} else {
				if len(item.Name())+width > int(tw) {
					entries += "\n"
					width = 0
				}

				entries += item.Name() + " "
				width += len(item.Name()) + len(" ")
			}

			if index+1 == len(items) {
				entries += "\n"
			}
		}

		fmt.Printf("%s", entries)
		return nil
	}

	if len(operands) < 1 {
		if err = list("."); err != nil {
			return err
		}
	} else {
		for _, operand = range operands {
			if err = ex.Exists(operand); err != nil {
				err = errors.New("ls: " + operand + ": No such file or directory")
				return err
			}

			if err = list(operand); err != nil {
				return err
			}
		}
	}

	return nil
}
