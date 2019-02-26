// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/groups"
	"github.com/jcmdln/flagger"
)

type groupsCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *groupsCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "groups", "[-bct] [FILE ...]"
	u.description = "Compute and check SHA1 message digest"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *groupsCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		if len(s) == 0 {
			if err := groups.Groups(""); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		if err := groups.Groups(data[0]); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	Command.Add("groups", &groupsCmd{})
}
