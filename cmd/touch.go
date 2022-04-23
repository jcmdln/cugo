// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/touch"
)

type touchCmd struct {
	touch.Options
}

func (u *touchCmd) Init() *flag.FlagSet {
	touch := flag.NewFlagSet("touch", flag.ExitOnError)
	touch.BoolVar(&u.Access, "a", false, "Change the access time")
	touch.BoolVar(&u.Create, "c", false, "Do not create missing files")
	touch.StringVar(&u.Date, "d", "",
		"Change access and modified time as RFC3339Nano")
	touch.BoolVar(&u.Modified, "m", false, "Change the modified time")
	touch.StringVar(&u.Reference, "r", "",
		"Use access and modified time of file")
	return touch
}

func (u *touchCmd) Run(s []string) error {
	if err := u.Touch(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["touch"] = &touchCmd{}
}
