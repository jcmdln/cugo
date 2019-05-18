// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/sha256sum"
	"github.com/jcmdln/flagger"
)

type sha256sumCmd struct {
	name        string
	usage       string
	description string

	help bool
	sha256sum.Options
}

func (u *sha256sumCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "sha256sum", "[-bct] [FILE ...]"
	u.description = "Compute and check SHA256 message digest"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *sha256sumCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		err  error
		data []string
	)

	if data, err = flags.Parse(s); err != nil {
		if err.Error() != "missing operand" {
			err = fmt.Errorf("%s: %s", u.name, err)
			return err
		}
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = u.Sha256sum(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("sha256sum", &sha256sumCmd{})
}
