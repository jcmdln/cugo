// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build testing

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/ls"
	"github.com/jcmdln/flagger"
)

type lsCmd struct {
	name        string
	usage       string
	description string

	help bool
	ls.Options
}

func (u *lsCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "ls", "[-aR] TARGET ..."
	u.description = "List directory contents"

	flags.BoolVar(&u.All, "Include directory entries that begin with '.'", "-a", "--all")
	flags.BoolVar(&u.Recursive, "Recursively traverse folders", "-R", "--recursive")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *lsCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		u.Ls([]string{"."})
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		u.Ls(data)
	}
	return nil
}

func init() {
	Command.Add("ls", &lsCmd{})
}
