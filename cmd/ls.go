// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/hlfstr/cugo/src/ls"
	"github.com/hlfstr/flagger"
)

type lsCmd struct{}

func (l *lsCmd) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&ls.All, "Include directory entries that begin with '.'", "-a", "--all")
	flags.BoolVar(&ls.Recursive, "Recursively traverse folders", "--recursive", "-r")
}

func (l *lsCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		ls.Ls(data)
	}
	return nil
}

func init() {
	Command.Add("ls", &lsCmd{})
}
