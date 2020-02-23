// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

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
			err = fmt.Errorf("%s: %s", u.name, err)
			return err
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
