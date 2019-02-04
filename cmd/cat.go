// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/cat"
	"github.com/jcmdln/flagger"
)

type catCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *catCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "cat", "[-u] [FILE ...]"
	u.description = "Concatenate and print files"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *catCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		cat.Cat(data)
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		cat.Cat(data)
	}

	return nil
}

func init() {
	Command.Add("cat", &catCmd{})
}
