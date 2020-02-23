// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/chmod"
	"github.com/jcmdln/flagger"
)

type chmodCmd struct {
	name        string
	usage       string
	description string

	help bool
	chmod.Options
}

func (u *chmodCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "chmod", "[-R] MODE FILE ..."
	u.description = "Change file modes"

	flags.BoolVar(&u.Recursive, "Change files and directories recursively", "-R", "--recursive")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *chmodCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		err  error
		data []string
	)

	if data, err = flags.Parse(s); err != nil {
		err = fmt.Errorf("%s: %s", u.name, err)
		return err
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = u.Chmod(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("chmod", &chmodCmd{})
}
