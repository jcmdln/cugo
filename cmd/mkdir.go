// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/mkdir"
)

type mkdirCmd struct {
	mkdir.Option
}

func (u *mkdirCmd) Init() *flag.FlagSet {
	mkdir := flag.NewFlagSet("mkdir", flag.ExitOnError)
	mkdir.UintVar(&u.Mode, "m", 0755, "Set permissions to MODE value")
	mkdir.BoolVar(&u.Parents, "p", false, "Create missing parent directories")
	mkdir.BoolVar(&u.Verbose, "v", false, "Display each created directory")
	return mkdir
}

func (u *mkdirCmd) Run(s []string) error {
	if _, err := u.Mkdir(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["mkdir"] = &mkdirCmd{}
}
