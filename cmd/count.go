// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/count"
	"github.com/jcmdln/flagger"
)

type countCmd struct{}

func (b *countCmd) Prepare(flags *flagger.Flags) {
	// This utility has no flags
}

func (b *countCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		count.Count(data)
	}
	return nil
}

func init() {
	Command.Add("count", &countCmd{})
}
