// uname.go
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

	"github.com/jcmdln/cugo/src/uname"
)

type unameCmd struct {
	uname.Options
}

func (u *unameCmd) Init() *flag.FlagSet {
	uname := flag.NewFlagSet("uname", flag.ExitOnError)
	uname.BoolVar(&u.All, "a", false,
		"Behave as though all options were specified")
	uname.BoolVar(&u.Machine, "m", false, "Print the machine hardware name")
	uname.BoolVar(&u.Nodename, "n", false, "Print the nodename / network name")
	uname.BoolVar(&u.Release, "r", false, "Print the operating system release")
	uname.BoolVar(&u.Sysname, "s", false, "Print the operating system name")
	uname.BoolVar(&u.Version, "v", false, "Print the operating system version")
	return uname
}

func (u *unameCmd) Run(s []string) error {
	if err := u.Uname(); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["uname"] = &unameCmd{}
}
