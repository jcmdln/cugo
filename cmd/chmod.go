// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/chmod"
)

type chmodCmd struct {
	chmod.Options
}

func (u *chmodCmd) Init() *flag.FlagSet {
	chmod := flag.NewFlagSet("chmod", flag.ExitOnError)
	chmod.BoolVar(&u.Recursive, "r", false,
		"Change files and directories recursively")
	return chmod
}

func (u *chmodCmd) Run(s []string) error {
	var (
		err error
	)

	if err = u.Chmod(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["chmod"] = &chmodCmd{}
}
