// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/false"
)

type falseCmd struct{}

func (u *falseCmd) Init() *flag.FlagSet {
	False := flag.NewFlagSet("false", flag.ExitOnError)
	return False
}

func (u *falseCmd) Run(s []string) error {
	false.False()
	return nil
}

func init() {
	Commands["false"] = &falseCmd{}
}
