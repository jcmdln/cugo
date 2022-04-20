// rmdir.go
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

	"github.com/jcmdln/cugo/src/rmdir"
)

type rmdirCmd struct {
	rmdir.Options
}

func (u *rmdirCmd) Init() *flag.FlagSet {
	rmdir := flag.NewFlagSet("rmdir", flag.ExitOnError)
	rmdir.BoolVar(&u.Parents, "p", false, "Remove parent directories")
	rmdir.BoolVar(&u.Verbose, "v", false,
		"Print a message when actions are taken")
	return rmdir
}

func (u *rmdirCmd) Run(s []string) error {
	if err := u.Rmdir(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["rmdir"] = &rmdirCmd{}
}
