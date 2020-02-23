// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/sha1sum"
	"github.com/jcmdln/flagger"
)

type sha1sumCmd struct {
	name        string
	usage       string
	description string

	help bool
	sha1sum.Options
}

func (u *sha1sumCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "sha1sum", "[-bct] [FILE ...]"
	u.description = "Compute and check SHA1 message digest"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *sha1sumCmd) Action(s []string, flags *flagger.Flags) error {
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

	if err = u.Sha1sum(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("sha1sum", &sha1sumCmd{})
}
