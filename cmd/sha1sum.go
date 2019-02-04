// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/sha1sum"
	"github.com/jcmdln/flagger"
)

type sha1sumCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *sha1sumCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "sha1sum", "[-bct] [FILE ...]"
	u.description = "Compute and check SHA1 message digest"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *sha1sumCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		sha1sum.Sha1sum(data)
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		sha1sum.Sha1sum(data)
	}

	return nil
}

func init() {
	Command.Add("sha1sum", &sha1sumCmd{})
}
