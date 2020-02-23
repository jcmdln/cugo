// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/true"
	"github.com/jcmdln/flagger"
)

type trueCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *trueCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "true", ""
	u.description = "Return true value"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *trueCmd) Action(s []string, flags *flagger.Flags) error {
	var err error
	if _, err = flags.Parse(s); err != nil {
		if err.Error() != "missing operand" {
			err = fmt.Errorf("%s: %s", u.name, err)
			return err
		}
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	true.True()

	return nil
}

func init() {
	Command.Add("true", &trueCmd{})
}
