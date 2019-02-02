// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/pwd"
	"github.com/jcmdln/flagger"
)

type pwdCmd struct{}

func (m *pwdCmd) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&pwd.L, "Read current dir from env, including symlinks", "-L")
	flags.BoolVar(&pwd.P, "Absolute path of current dir without symlinks", "-P")
}

func (m *pwdCmd) Action(s []string, flags *flagger.Flags) error {
	pwd.Pwd()

	return nil
}

func init() {
	Command.Add("pwd", &pwdCmd{})
}
