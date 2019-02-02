// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/whoami"
	"github.com/jcmdln/flagger"
)

type whoamiCmd struct{}

func (m *whoamiCmd) Prepare(flags *flagger.Flags) {
	// This utility has no flags
}

func (m *whoamiCmd) Action(s []string, flags *flagger.Flags) error {
	whoami.Whoami()

	return nil
}

func init() {
	Command.Add("whoami", &whoamiCmd{})
}
