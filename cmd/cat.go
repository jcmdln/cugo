// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/cat"
)

type catCmd struct {
	cat.Options
}

func (u *catCmd) Init() *flag.FlagSet {
	cat := flag.NewFlagSet("cat", flag.ExitOnError)
	cat.BoolVar(&u.Unbuffered, "u", false, "Unbuffered output")
	return cat
}

func (u *catCmd) Run(s []string) error {
	if err := u.Cat(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["cat"] = &catCmd{}
}
