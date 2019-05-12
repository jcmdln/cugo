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
	u.description = ""

	flags.BoolVar(&u.help, "show group memberships", "-h", "--help")
}

func (u *groupsCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		user string
	)

	users, _ := flags.Parse(s)
	if len(users) < 1 {
		user = ""
	} else {
		user = users[0]
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err := groups.Groups(user); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("groups", &groupsCmd{})
}
