// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/head"
)

type headCmd struct {
	head.Option
}

func (u *headCmd) Init() *flag.FlagSet {
	head := flag.NewFlagSet("head", flag.ExitOnError)
	head.IntVar(&u.Number, "n", 0,
		"Copy the first NUMBER of lines of each file")
	return head
}

func (u *headCmd) Run(s []string) error {
	if err := u.Head(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["head"] = &headCmd{}
}
