// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/ls"
	"github.com/jcmdln/flagger"
)

type lsCmd struct {
	name        string
	usage       string
	description string

	help bool
	ls.Options
}

func (u *lsCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "ls", "[-aR] TARGET ..."
	u.description = "List directory contents"

	flags.BoolVar(&u.All, "Show all entries other than '.' and '..'", "-A")
	flags.BoolVar(&u.Recursive, "Recursively list directories and their entries", "-R")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *lsCmd) Action(s []string, flags *flagger.Flags) error {
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

	if err = u.Ls(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("ls", &lsCmd{})
}
