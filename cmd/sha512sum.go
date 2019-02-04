// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/sha512sum"
	"github.com/jcmdln/flagger"
)

type sha512sumCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *sha512sumCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "sha512sum", "[-bct] [FILE ...]"
	u.description = "Compute and check SHA512 message digest"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *sha512sumCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		sha512sum.Sha512sum(data)
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		sha512sum.Sha512sum(data)
	}

	return nil
}

func init() {
	Command.Add("sha512sum", &sha512sumCmd{})
}
