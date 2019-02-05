// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/touch"
	"github.com/jcmdln/flagger"
)

type touchCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *touchCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "touch", "[-acm] [-d ccyy-mm-ddTHH:MM:SS[.frac][Z]] [-r FILE] FILE ..."
	u.description = "Change file access and modification times"

	flags.BoolVar(&touch.Access, "Change the access time", "-a")
	flags.BoolVar(&touch.Create, "Do not create missing files", "-c")
	flags.StringVar(&touch.Date, "", "Change access and modified time as RFC3339Nano", "-d")
	flags.BoolVar(&touch.Modified, "Change the modified time", "-m")
	flags.StringVar(&touch.Reference, "", "Use access and modified time of file", "-r")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *touchCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		touch.Touch(data)
	} else {
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		touch.Touch(data)
	}

	return nil
}

func init() {
	Command.Add("touch", &touchCmd{})
}
