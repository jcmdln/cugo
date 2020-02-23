// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/yes"
	"github.com/jcmdln/flagger"
)

type yesCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *yesCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "yes", ""
	u.description = "Be repetitively affirmative"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *yesCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		yes.Yes([]string{})
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		yes.Yes(data)
	}

	return nil
}

func init() {
	Command.Add("yes", &yesCmd{})
}
