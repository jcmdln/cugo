// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/sha256sum"
)

type sha256sumCmd struct {
	sha256sum.Options
}

func (u *sha256sumCmd) Init() *flag.FlagSet {
	sha256sum := flag.NewFlagSet("sha256sum", flag.ExitOnError)
	return sha256sum
}

func (u *sha256sumCmd) Run(s []string) error {
	if err := u.Sha256sum(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["sha256sum"] = &sha256sumCmd{}
}
