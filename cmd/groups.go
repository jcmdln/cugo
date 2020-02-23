// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package cmd

import (
	"fmt"

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
		data []string
		err  error
		user string
	)

	if data, err = flags.Parse(s); err != nil {
		if err.Error() != "missing operand" {
			err = fmt.Errorf("%s: %s", u.name, err)
			return err
		}
	}

	if len(data) < 1 {
		user = ""
	} else {
		user = data[0]
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = groups.Groups(user); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("groups", &groupsCmd{})
}
