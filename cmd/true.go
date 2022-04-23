// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/true"
)

type trueCmd struct{}

func (u *trueCmd) Init() *flag.FlagSet {
	true := flag.NewFlagSet("true", flag.ExitOnError)
	return true
}

func (u *trueCmd) Run(s []string) error {
	true.True()
	return nil
}

func init() {
	Commands["true"] = &trueCmd{}
}
