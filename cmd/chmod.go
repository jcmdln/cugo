// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/chmod"
	"github.com/jcmdln/flagger"
)

type chmodCmd struct{}

func (b *chmodCmd) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&chmod.Recursive, "Change files and directories recursively", "-R", "--recursive")
}

func (b *chmodCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		chmod.Chmod(data)
	}

	return nil
}

func init() {
	Command.Add("chmod", &chmodCmd{})
}
