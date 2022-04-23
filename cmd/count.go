// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/count"
)

type countCmd struct{}

func (u *countCmd) Init() *flag.FlagSet {
	count := flag.NewFlagSet("count", flag.ExitOnError)
	return count
}

func (u *countCmd) Run(s []string) error {
	if err := count.Count(s); err != nil {
		return err
	}
	return nil
}

func init() {
	Commands["count"] = &countCmd{}
}
