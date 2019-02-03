// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/hostname"
	"github.com/jcmdln/flagger"
)

type hostnameCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *hostnameCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "hostname", "[-s] HOSTNAME"
	u.description = "Set or print name of current host system"

	flags.BoolVar(&hostname.Strip, "Trim any domain information from the printed name", "-s", "--strip")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *hostnameCmd) Action(s []string, flags *flagger.Flags) error {
	var d string

	lenCheck := func(data []string) {
		if len(data) > 1 {
			help.Help(u.name, u.usage, u.description, flags)
		} else if len(data) == 0 {
			d = ""
		} else {
			d = data[0]
		}
	}

	if data, err := flags.Parse(s); err != nil {
		lenCheck(data)
		hostname.Hostname(d)
	} else {
		lenCheck(data)
		if u.help {
			help.Help(u.name, u.usage, u.description, flags)
		}

		hostname.Hostname(d)
	}

	return nil
}

func init() {
	Command.Add("hostname", &hostnameCmd{})
}
