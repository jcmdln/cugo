// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/rmdir"
	"github.com/jcmdln/flagger"
)

type rmdirCmd struct{}

func (m *rmdirCmd) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&rmdir.Parents, "Remove parent directories", "-p", "--parents")
	flags.BoolVar(&rmdir.Verbose, "Print a message when actions are taken", "-v", "--verbose")
}

func (m *rmdirCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		rmdir.Rmdir(data)
	}
	return nil
}

func init() {
	Command.Add("rmdir", &rmdirCmd{})
}
