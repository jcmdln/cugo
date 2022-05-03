// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/jcmdln/cugo/src/chmod"
)

type chmodCmd struct {
	chmod.Option
}

func (u *chmodCmd) Init() *flag.FlagSet {
	chmod := flag.NewFlagSet("chmod", flag.ExitOnError)
	chmod.BoolVar(&u.Recursive, "r", false, "Change files and directories recursively")
	return chmod
}

func (u *chmodCmd) Run(s []string) error {
	var (
		err   error
		file  *os.File
		files []*os.File
	)

	if len(s) < 2 {
		return fmt.Errorf("chmod: missing required argument MODE")
	}

	for _, f := range s[1:] {
		if file, err = os.Open(f); err != nil {
			return err
		}
		files = append(files, file)
	}

	if err = u.Chmod(s[0], files); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["chmod"] = &chmodCmd{}
}
