// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/false"
	"github.com/jcmdln/flagger"
)

type falseCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *falseCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "false", ""
	u.description = "Return false value"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *falseCmd) Action(s []string, flags *flagger.Flags) error {
	if _, err := flags.Parse(s); err != nil {
		false.False()
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}
	}

	return nil
}

func init() {
	Command.Add("false", &falseCmd{})
}
