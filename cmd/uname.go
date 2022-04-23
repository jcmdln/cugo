// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/uname"
)

type unameCmd struct {
	uname.Option
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
