// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/hostname"
	"github.com/jcmdln/flagger"
)

type hostnameCmd struct{}

func (l *hostnameCmd) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&hostname.Strip, "Trim any domain information from the printed name", "-s", "--strip")
}

func (l *hostnameCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		hostname.Hostname(data)
	} else {
		hostname.Hostname(data)
	}

	return nil
}

func init() {
	Command.Add("hostname", &hostnameCmd{})
}
