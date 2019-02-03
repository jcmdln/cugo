// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/true"
	"github.com/jcmdln/flagger"
)

type trueCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *trueCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "true", ""
	u.description = "Return true value"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *trueCmd) Action(s []string, flags *flagger.Flags) error {
	if _, err := flags.Parse(s); err != nil {
		true.True()
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}
	}

	return nil
}

func init() {
	Command.Add("true", &trueCmd{})
}
