// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/cat"
	"github.com/jcmdln/flagger"
)

type catCmd struct {
	name        string
	usage       string
	description string

	help bool
	cat.Options
}

func (u *catCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "cat", "[-u] [FILE ...]"
	u.description = "Concatenate and print files"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
	flags.BoolVar(&u.Unbuffered, "Unbuffered output", "-u")
}

func (u *catCmd) Action(s []string, flags *flagger.Flags) error {
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

	if err = u.Cat(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("cat", &catCmd{})
}
