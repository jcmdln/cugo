// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/rmdir"
	"github.com/jcmdln/flagger"
)

type rmdirCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *rmdirCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "rmdir", "[-pv] DIRECTORY ..."
	u.description = "Remove directories"

	flags.BoolVar(&rmdir.Parents, "Remove parent directories", "-p", "--parents")
	flags.BoolVar(&rmdir.Verbose, "Print a message when actions are taken", "-v", "--verbose")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *rmdirCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		rmdir.Rmdir(data)
	}

	return nil
}

func init() {
	Command.Add("rmdir", &rmdirCmd{})
}
