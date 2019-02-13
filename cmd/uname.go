// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
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
	if _, err := flags.Parse(s); err != nil {
		u.Uname()
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		u.Uname()
	}

	return nil
}

func init() {
	Command.Add("uname", &unameCmd{})
}
