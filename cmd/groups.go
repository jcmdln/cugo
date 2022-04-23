// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/groups"
)

type groupsCmd struct{}

func (u *groupsCmd) Init() *flag.FlagSet {
	groups := flag.NewFlagSet("groups", flag.ExitOnError)
	return groups
}

func (u *groupsCmd) Run(s []string) error {
	var user string

	if len(s) < 1 {
		user = ""
	} else {
		user = s[0]
	}

	if err := groups.Groups(user); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["groups"] = &groupsCmd{}
}
