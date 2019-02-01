// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/hlfstr/cugo/src/mkdir"
	"github.com/hlfstr/flagger"
)

type mkdirCmd struct{}

func (m *mkdirCmd) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&mkdir.Parents, "Create missing parent directories", "-p", "--parents")
	flags.BoolVar(&mkdir.Verbose, "Display each directory after it was created", "-v", "--verbose")
	mkdir.Mode = 0755 //Unit32 is not part of flagger yet
}

func (m *mkdirCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		mkdir.Mkdir(data)
	}
	return nil
}

func init() {
	Command.Add("mkdir", &mkdirCmd{})
}
