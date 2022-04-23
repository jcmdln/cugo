// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/sha512sum"
)

type sha512sumCmd struct {
	sha512sum.Options
}

func (u *sha512sumCmd) Init() *flag.FlagSet {
	sha512sum := flag.NewFlagSet("sha512sum", flag.ExitOnError)
	return sha512sum
}

func (u *sha512sumCmd) Run(s []string) error {
	if err := u.Sha512sum(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["sha512sum"] = &sha512sumCmd{}
}
