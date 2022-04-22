// basename.go
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
