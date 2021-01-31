// sha384sum.go
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
	"github.com/jcmdln/cugo/src/sha384sum"
	"github.com/jcmdln/flagger"
)

type sha384sumCmd struct {
	name        string
	usage       string
	description string

	help bool
	sha384sum.Options
}

func (u *sha384sumCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "sha384sum", "[-bct] [FILE ...]"
	u.description = "Compute and check SHA384 message digest"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *sha384sumCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		err  error
		data []string
	)

	if data, err = flags.Parse(s); err != nil {
		if err.Error() != "missing operand" {
			err = fmt.Errorf("%s: %s", u.name, err)
			return err
		}
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = u.Sha384sum(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("sha384sum", &sha384sumCmd{})
}
