// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"
	"fmt"

	"github.com/jcmdln/cugo/src/basename"
)

type basenameCmd struct{}

func (u *basenameCmd) Init() *flag.FlagSet {
	basename := flag.NewFlagSet("basename", flag.ExitOnError)
	return basename
}

func (u *basenameCmd) Run(s []string) error {
	var (
		err     error
		operand string
		suffix  string
		result  string
	)

	dataLen := len(s)
	switch {
	case dataLen < 1:
		err = fmt.Errorf("basename: missing operand")
		return err
	case dataLen > 2:
		err = fmt.Errorf("basename: extra operands")
		return err
	case dataLen == 2:
		operand = s[0]
		suffix = s[1]
	default:
		operand = s[0]
		suffix = ""
	}

	if result, err = basename.Basename(operand, suffix); err != nil {
		return err
	}

	if _, err = fmt.Printf("%s\n", result); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["basename"] = &basenameCmd{}
}
