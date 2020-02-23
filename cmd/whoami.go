// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/whoami"
	"github.com/jcmdln/flagger"
)

type whoamiCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *whoamiCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "whoami", ""
	u.description = "Display effective user ID"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *whoamiCmd) Action(s []string, flags *flagger.Flags) error {
	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if len(s) > 0 {
		t := strings.Join(s, " ")
		fmt.Printf("cugo: whoami: extra operand '%s'\n", t)
	} else {
		whoami.Whoami()
	}

	return nil
}

func init() {
	Command.Add("whoami", &whoamiCmd{})
}
