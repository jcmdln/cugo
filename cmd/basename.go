// Copyright 2021 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that may
// be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/basename"
	"github.com/jcmdln/flagger"
)

type basenameCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *basenameCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "basename", "STRING [SUFFIX]"
	u.description = "Return non-directory portion of a pathname"
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *basenameCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		data    []string
		err     error
		operand string
		suffix  string
		ret     string
	)

	if data, err = flags.Parse(s); err != nil {
		err = fmt.Errorf("%s: %s", u.name, err)
		return err
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	dataLen := len(data)
	switch {
	case dataLen < 1:
		err = fmt.Errorf("%s: missing operand", u.name)
		return err
	case dataLen > 2:
		err = fmt.Errorf("%s: extra operands", u.name)
		return err
	case dataLen == 2:
		operand = data[0]
		suffix = data[1]
	default:
		operand = data[0]
		suffix = ""
	}

	if ret, err = basename.Basename(operand, suffix); err != nil {
		return err
	}

	if _, err = fmt.Printf("%s\n", ret); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("basename", &basenameCmd{})
}
