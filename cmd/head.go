// head.go
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

	"github.com/jcmdln/cugo/src/head"
)

type headCmd struct {
	head.Options
}

func (u *headCmd) Init() *flag.FlagSet {
	head := flag.NewFlagSet("head", flag.ExitOnError)
	head.IntVar(&u.Number, "n", 0,
		"Copy the first NUMBER of lines of each file")
	return head
}

func (u *headCmd) Run(s []string) error {
	if err := u.Head(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["head"] = &headCmd{}
}
