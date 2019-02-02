// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/touch"
	"github.com/jcmdln/flagger"
)

type touchCmd struct{}

func (m *touchCmd) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&touch.Access, "Change the access time", "-a")
	flags.BoolVar(&touch.Create, "Do not create missing files", "-c")
	flags.StringVar(&touch.Date, "Change access and modified time as per ISO8601/RFC3339Nano", "-d")
	flags.BoolVar(&touch.Modified, "Change the modified time", "-m")
	flags.StringVar(&touch.Reference, "Use access and modified time from reference file", "-r", "--verbose")
}

func (m *touchCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		touch.Touch(data)
	}
	return nil
}

func init() {
	Command.Add("touch", &touchCmd{})
}
