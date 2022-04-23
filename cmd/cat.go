// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"
	"os"

	"github.com/jcmdln/cugo/src/cat"
)

type catCmd struct {
	cat.Option
}

func (u *catCmd) Init() *flag.FlagSet {
	cat := flag.NewFlagSet("cat", flag.ExitOnError)
	cat.BoolVar(&u.Unbuffered, "u", false, "Unbuffered output")
	return cat
}

func (u *catCmd) Run(s []string) error {
	var (
		err   error
		file  *os.File
		files []*os.File
	)

	if len(files) == 0 {
		files = append(files, os.Stdin)
	}

	for _, f := range s {
		if file, err = os.Open(f); err != nil {
			return err
		}
		files = append(files, file)
	}

	if err := u.Cat(files); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["cat"] = &catCmd{}
}
