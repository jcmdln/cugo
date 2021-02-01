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
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/uname"
	"github.com/jcmdln/flagger"
)

type unameCmd struct {
	name        string
	usage       string
	description string

	help bool
	uname.Options
}

func (u *unameCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "uname", "[-amnprsv]"
	u.description = "print operating system name"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
	flags.BoolVar(&u.All, "Behave as though all options were specified", "-a")
	flags.BoolVar(&u.Machine, "Print the machine hardware name", "-m")
	flags.BoolVar(&u.Nodename, "Print the nodename (aka network name)", "-n")
	flags.BoolVar(&u.Release, "Print the operating system release", "-r")
	flags.BoolVar(&u.Sysname, "Print the operating system name", "-s")
	flags.BoolVar(&u.Version, "Print the operating system version", "-v")
}

func (u *unameCmd) Action(s []string, flags *flagger.Flags) error {
	var err error

	if _, err = flags.Parse(s); err != nil {
		if err.Error() != "missing operand" {
			return fmt.Errorf("%s: %s", u.name, err)
		}
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = u.Uname(); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("uname", &unameCmd{})
}
