// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/sha224sum"
)

type sha224sumCmd struct {
	sha224sum.Option
}

func (u *sha224sumCmd) Init() *flag.FlagSet {
	sha224sum := flag.NewFlagSet("sha224sum", flag.ExitOnError)
	return sha224sum
}

func (u *sha224sumCmd) Run(s []string) error {
	if err := u.Sha224sum(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["sha224sum"] = &sha224sumCmd{}
}
