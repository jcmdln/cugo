// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/rm"
)

type rmCmd struct {
	rm.Options
}

func (u *rmCmd) Init() *flag.FlagSet {
	rm := flag.NewFlagSet("rm", flag.ExitOnError)
	rm.BoolVar(&u.Dir, "d", false, "Remove empty directories")
	rm.BoolVar(&u.Force, "f", false, "Skip prompts and ignore warnings")
	rm.BoolVar(&u.Interactive, "i", false, "Prompt before each removal")
	rm.BoolVar(&u.Recursive, "r", false,
		"Remove directories and their contents recursively")
	rm.BoolVar(&u.Verbose, "v", false,
		"Print a message when actions are taken")
	return rm
}

func (u *rmCmd) Run(s []string) error {
	if err := u.Rm(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["rm"] = &rmCmd{}
}
