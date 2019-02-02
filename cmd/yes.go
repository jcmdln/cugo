// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/flagger"
	"github.com/jcmdln/cugo/src/yes"
)

type yesCmd struct{}

func (b *yesCmd) Prepare(flags *flagger.Flags) {
	// This utility has no flags
}

func (b *yesCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		yes.Yes([]string{})
	} else {
		yes.Yes(data)
	}

	return nil
}

func init() {
	Command.Add("yes", &yesCmd{})
}
