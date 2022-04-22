// ls.go
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

//go:build testing
// +build testing

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
