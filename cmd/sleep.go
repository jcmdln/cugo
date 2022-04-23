// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/sleep"
)

type sleepCmd struct{}

func (u *sleepCmd) Init() *flag.FlagSet {
	sleep := flag.NewFlagSet("sleep", flag.ExitOnError)
	return sleep
}

func (u *sleepCmd) Run(s []string) error {
	if err := sleep.Sleep(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["sleep"] = &sleepCmd{}
}
