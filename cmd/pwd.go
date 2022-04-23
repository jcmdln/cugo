// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/pwd"
)

type pwdCmd struct {
	pwd.Option
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
