// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/rm"
	"github.com/jcmdln/flagger"
)

type rmCmd struct {
	name        string
	usage       string
	description string

	help bool
	rm.Options
}

func (u *rmCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "rm", "[-dfiPRrv] FILE ..."
	u.description = "Remove directory entries"

	flags.BoolVar(&u.Dir, "Remove empty directories", "-d")
	flags.BoolVar(&u.Force, "Skip prompts and ignore warnings", "-f")
	flags.BoolVar(&u.Interactive, "Prompt before each removal", "-i")
	flags.BoolVar(&u.Recursive, "Remove directories and their contents recursively", "-r", "-R")
	flags.BoolVar(&u.Verbose, "Print a message when actions are taken", "-v")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *rmCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		data []string
		err  error
	)

	if data, err = flags.Parse(s); err != nil {
		err = fmt.Errorf("%s: %s", u.name, err)
		return err
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = u.Rm(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("rm", &rmCmd{})
}
