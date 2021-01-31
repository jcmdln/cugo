// hostname.go
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
	"github.com/jcmdln/cugo/src/hostname"
	"github.com/jcmdln/flagger"
)

type hostnameCmd struct {
	name        string
	usage       string
	description string

	help bool
	hostname.Options
}

func (u *hostnameCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "hostname", "[-s] HOSTNAME"
	u.description = "Set or print name of current host system"

	flags.BoolVar(&u.Strip, "Trim any domain information from the printed name", "-s", "--strip")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *hostnameCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		data     []string
		err      error
		hostname string
	)

	if data, err = flags.Parse(s); err != nil {
		if err.Error() != "missing operand" {
			err = fmt.Errorf("%s: %s", u.name, err)
			return err
		}
	}

	if len(data) == 0 {
		hostname = ""
	} else {
		hostname = data[0]
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err = u.Hostname(hostname); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("hostname", &hostnameCmd{})
}
