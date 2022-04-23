// SPDX-License-Identifier: ISC

//go:build testing

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/src/ls"
)

type lsCmd struct {
	ls.Option
}

func (u *lsCmd) Init() *flag.FlagSet {
	ls := flag.NewFlagSet("ls", flag.ExitOnError)
	ls.BoolVar(&u.All,
		"Show all entries other than '.' and '..'", "-A")
	ls.BoolVar(&u.Recursive, "R",
		"Recursively list directories and their entries")
	return ls
}

func (u *lsCmd) Run(s []string) error {
	if err := u.Ls(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("ls", &lsCmd{})
}
