// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/hostname"
)

type hostnameCmd struct {
	hostname.Option
}

func (u *hostnameCmd) Init() *flag.FlagSet {
	hostname := flag.NewFlagSet("hostname", flag.ExitOnError)
	hostname.BoolVar(&u.Strip, "s", false,
		"Trim any domain information from the printed name")
	return hostname
}

func (u *hostnameCmd) Run(s []string) error {
	var hostname string

	if len(s) == 0 {
		hostname = ""
	} else {
		hostname = s[0]
	}

	if err := u.Hostname(hostname); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["hostname"] = &hostnameCmd{}
}
