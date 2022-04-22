// basename.go
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

package cmd

import (
	"flag"
	"fmt"

	"github.com/jcmdln/cugo/src/basename"
)

type basenameCmd struct{}

func (u *basenameCmd) Init() *flag.FlagSet {
	basename := flag.NewFlagSet("basename", flag.ExitOnError)
	return basename
}

func (u *basenameCmd) Run(s []string) error {
	var (
		err     error
		operand string
		suffix  string
		result  string
	)

	dataLen := len(s)
	switch {
	case dataLen < 1:
		err = fmt.Errorf("basename: missing operand")
		return err
	case dataLen > 2:
		err = fmt.Errorf("basename: extra operands")
		return err
	case dataLen == 2:
		operand = s[0]
		suffix = s[1]
	default:
		operand = s[0]
		suffix = ""
	}

	if result, err = basename.Basename(operand, suffix); err != nil {
		return err
	}

	if _, err = fmt.Printf("%s\n", result); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["basename"] = &basenameCmd{}
}
