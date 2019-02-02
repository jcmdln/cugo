// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/rm"
	"github.com/jcmdln/flagger"
)

type rmCmd struct{}

func (m *rmCmd) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&rm.Force, "Skip prompts and ignore warnings", "-f", "--force")
	flags.BoolVar(&rm.Dir, "Remove empty directories", "-d", "--dir")
	flags.BoolVar(&rm.Interactive, "Prompt before each removal", "-i")
	flags.BoolVar(&rm.Recursive, "Remove directories and their contents recursively", "-r", "-R", "--recursive")
	flags.BoolVar(&rm.Verbose, "Print a message when actions are taken", "-v", "--verbose")
}

func (m *rmCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		rm.Rm(data)
	}
	return nil
}

func init() {
	Command.Add("rm", &rmCmd{})
}
