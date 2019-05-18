// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/pwd"
	"github.com/jcmdln/flagger"
)

type pwdCmd struct {
	name        string
	usage       string
	description string

	help bool
	pwd.Options
}

func (u *pwdCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "pwd", "[-LP]"
	u.description = "return working directory name"

	flags.BoolVar(&u.L, "Read current dir from env, including symlinks", "-L")
	flags.BoolVar(&u.P, "Absolute path of current dir without symlinks", "-P")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *pwdCmd) Action(s []string, flags *flagger.Flags) error {
	var err error

	if _, err = flags.Parse(s); err != nil {
		if err.Error() != "missing operand" {
			err = fmt.Errorf("%s: %s", u.name, err)
			return err
		}
	}

	if len(s) > 0 {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = u.Pwd(); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("pwd", &pwdCmd{})
}
