// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/false"
	"github.com/jcmdln/flagger"
)

type falseCmd struct{}

func (m *falseCmd) Prepare(flags *flagger.Flags) {
	// This utility has no flags
}

func (m *falseCmd) Action(s []string, flags *flagger.Flags) error {
	false.False()

	return nil
}

func init() {
	Command.Add("false", &falseCmd{})
}
