// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/true"
	"github.com/jcmdln/flagger"
)

type trueCmd struct{}

func (m *trueCmd) Prepare(flags *flagger.Flags) {
	// This utility has no flags
}

func (m *trueCmd) Action(s []string, flags *flagger.Flags) error {
	true.True()

	return nil
}

func init() {
	Command.Add("true", &trueCmd{})
}
