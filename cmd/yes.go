// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/yes"
)

type yesCmd struct{}

func (u *yesCmd) Init() *flag.FlagSet {
	yes := flag.NewFlagSet("yes", flag.ExitOnError)
	return yes
}

func (u *yesCmd) Run(s []string) error {
	yes.Yes(s)
	return nil
}

func init() {
	Commands["yes"] = &yesCmd{}
}
