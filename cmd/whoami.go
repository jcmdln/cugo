// SPDX-License-Identifier: ISC

package cmd

import (
	"flag"
	"fmt"
	"strings"

	"github.com/jcmdln/cugo/src/whoami"
)

type whoamiCmd struct{}

func (u *whoamiCmd) Init() *flag.FlagSet {
	whoami := flag.NewFlagSet("whoami", flag.ExitOnError)
	return whoami
}

func (u *whoamiCmd) Run(s []string) error {
	if len(s) > 0 {
		t := strings.Join(s, " ")
		fmt.Printf("cugo: whoami: extra operand '%s'\n", t)
	} else {
		whoami.Whoami()
	}

	return nil
}

func init() {
	Commands["whoami"] = &whoamiCmd{}
}
