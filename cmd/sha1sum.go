// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/sha1sum"
)

type sha1sumCmd struct {
	sha1sum.Option
}

func (u *sha1sumCmd) Init() *flag.FlagSet {
	sha1sum := flag.NewFlagSet("sha1sum", flag.ExitOnError)
	return sha1sum
}

func (u *sha1sumCmd) Run(s []string) error {
	if err := u.Sha1sum(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["sha1sum"] = &sha1sumCmd{}
}
