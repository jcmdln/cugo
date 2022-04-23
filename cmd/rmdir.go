// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/rmdir"
)

type rmdirCmd struct {
	rmdir.Option
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
