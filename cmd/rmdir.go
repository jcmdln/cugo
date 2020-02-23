// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/rmdir"
	"github.com/jcmdln/flagger"
)

type rmdirCmd struct {
	name        string
	usage       string
	description string

	help bool
	rmdir.Options
}

func (u *rmdirCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "rmdir", "[-pv] DIRECTORY ..."
	u.description = "Remove directories"

	flags.BoolVar(&u.Parents, "Remove parent directories", "-p", "--parents")
	flags.BoolVar(&u.Verbose, "Print a message when actions are taken", "-v", "--verbose")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *rmdirCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		data []string
		err  error
	)

	if data, err = flags.Parse(s); err != nil {
		err = fmt.Errorf("%s: %s", u.name, err)
		return err
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = u.Rmdir(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("rmdir", &rmdirCmd{})
}
