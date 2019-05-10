// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/basename"
	"github.com/jcmdln/flagger"
)

type basenameCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *basenameCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "basename", "STRING [SUFFIX]"
	u.description = "Return non-directory portion of a pathname"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *basenameCmd) Action(s []string, flags *flagger.Flags) error {
	data, err := flags.Parse(s)
	if err != nil {
		return err
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err := basename.Basename(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("basename", &basenameCmd{})
}
