// sleep.go
//
// Copyright 2021 Johnathan C. Maudlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/sleep"
	"github.com/jcmdln/flagger"
)

type sleepCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *sleepCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "sleep", "DURATION ..."
	u.description = "Suspend execution for an interval of time"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *sleepCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		data []string
		err  error
	)

	if data, err = flags.Parse(s); err != nil {
		return fmt.Errorf("%s: %s", u.name, err)
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = sleep.Sleep(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("sleep", &sleepCmd{})
}
