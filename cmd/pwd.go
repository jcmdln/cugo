// pwd.go
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

	"github.com/jcmdln/cugo/src/pwd"
)

type pwdCmd struct {
	pwd.Options
}

func (u *pwdCmd) Init() *flag.FlagSet {
	pwd := flag.NewFlagSet("pwd", flag.ExitOnError)
	pwd.BoolVar(&u.L, "L", false,
		"Read current dir from env, including symlinks")
	pwd.BoolVar(&u.P, "P", false,
		"Absolute path of current dir without symlinks")
	return pwd
}

func (u *pwdCmd) Run(s []string) error {
	if err := u.Pwd(); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["pwd"] = &pwdCmd{}
}
