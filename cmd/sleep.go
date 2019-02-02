// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"strings"

	"github.com/jcmdln/cugo/src/sleep"
	"github.com/jcmdln/flagger"
)

type sleepCmd struct{}

func (m *sleepCmd) Prepare(flags *flagger.Flags) {
	// This utility has no flags
}

func (m *sleepCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		sleep.Sleep(strings.Join(data, " "))
	}
	return nil
}

func init() {
	Command.Add("sleep", &sleepCmd{})
}
