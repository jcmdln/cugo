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

//go:build testing

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/src/ls"
)

type lsCmd struct {
	ls.Options
}

func (u *lsCmd) Init() *flag.FlagSet {
	ls := flag.NewFlagSet("ls", flag.ExitOnError)
	ls.BoolVar(&u.All,
		"Show all entries other than '.' and '..'", "-A")
	ls.BoolVar(&u.Recursive, "R",
		"Recursively list directories and their entries")
	return ls
}

func (u *lsCmd) Run(s []string) error {
	if err := u.Ls(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("ls", &lsCmd{})
}
