// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/rm"
	"github.com/jcmdln/flagger"
)

type rmCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *rmCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "rm", "[-dfiPRrv] FILE ..."
	u.description = "Remove directory entries"

	flags.BoolVar(&rm.Dir, "Remove empty directories", "-d")
	flags.BoolVar(&rm.Force, "Skip prompts and ignore warnings", "-f")
	flags.BoolVar(&rm.Interactive, "Prompt before each removal", "-i")
	flags.BoolVar(&rm.Recursive, "Remove directories and their contents recursively", "-r", "-R")
	flags.BoolVar(&rm.Verbose, "Print a message when actions are taken", "-v")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *rmCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		rm.Rm(data)
	}
	return nil
}

func init() {
	Command.Add("rm", &rmCmd{})
}
