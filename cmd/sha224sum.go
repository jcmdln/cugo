// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/sha224sum"
	"github.com/jcmdln/flagger"
)

type sha224sumCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *sha224sumCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "sha224sum", "[-bct] [FILE ...]"
	u.description = "Compute and check SHA224 message digest"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *sha224sumCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		sha224sum.Sha224sum(data)
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		sha224sum.Sha224sum(data)
	}

	return nil
}

func init() {
	Command.Add("sha224sum", &sha224sumCmd{})
}
