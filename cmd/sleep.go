// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/sleep"
	"github.com/jcmdln/flagger"
)

type sleepCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *sleepCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "sleep", "DURATION ..."
	u.description = "Suspend execution for an interval of time"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *sleepCmd) Action(s []string, flags *flagger.Flags) error {
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

	if err = sleep.Sleep(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("sleep", &sleepCmd{})
}
