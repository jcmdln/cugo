// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/flagger"
	"github.com/jcmdln/cugo/src/basename"
)

type basenameCmd struct{}

func (b *basenameCmd) Prepare(flags *flagger.Flags) {
	// This utility has no flags
}

func (b *basenameCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		basename.Basename(data)
	}

	return nil
}

func init() {
	Command.Add("basename", &basenameCmd{})
}
