// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/sha384sum"
)

type sha384sumCmd struct {
	sha384sum.Options
}

func (u *sha384sumCmd) Init() *flag.FlagSet {
	sha384sum := flag.NewFlagSet("sha384sum", flag.ExitOnError)
	return sha384sum
}

func (u *sha384sumCmd) Run(s []string) error {
	if err := u.Sha384sum(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["sha384sum"] = &sha384sumCmd{}
}
