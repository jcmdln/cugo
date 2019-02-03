// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/chmod"
	"github.com/jcmdln/flagger"
)

type chmodCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *chmodCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "chmod", "[-R] MODE FILE ..."
	u.description = "Change file modes"

	flags.BoolVar(&chmod.Recursive, "Change files and directories recursively", "-R", "--recursive")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *chmodCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		chmod.Chmod(data)
	}

	return nil
}

func init() {
	Command.Add("chmod", &chmodCmd{})
}
