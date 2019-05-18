// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/hostname"
	"github.com/jcmdln/flagger"
)

type hostnameCmd struct {
	name        string
	usage       string
	description string

	help bool
	hostname.Options
}

func (u *hostnameCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "hostname", "[-s] HOSTNAME"
	u.description = "Set or print name of current host system"

	flags.BoolVar(&u.Strip, "Trim any domain information from the printed name", "-s", "--strip")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *hostnameCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		data     []string
		err      error
		hostname string
	)

	if data, err = flags.Parse(s); err != nil {
		if err.Error() != "missing operand" {
			err = fmt.Errorf("%s: %s", u.name, err)
			return err
		}
	}

	if len(data) == 0 {
		hostname = ""
	} else {
		hostname = data[0]
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = u.Hostname(hostname); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("hostname", &hostnameCmd{})
}
