// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/head"
	"github.com/jcmdln/flagger"
)

type headCmd struct {
	name        string
	usage       string
	description string

	help bool
	head.Options
}

func (u *headCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "head", "[-s] HEAD"
	u.description = "Set or print name of current host system"

	flags.IntVar(&u.Number, 10, "Copy the first NUMBER of lines of each file", "-n")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *headCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		if err := u.Head(data); err != nil {
			return err
		}
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		if err := u.Head(data); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	Command.Add("head", &headCmd{})
}
