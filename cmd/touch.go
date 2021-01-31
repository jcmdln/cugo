// touch.go
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
	"github.com/jcmdln/cugo/src/touch"
	"github.com/jcmdln/flagger"
)

type touchCmd struct {
	name        string
	usage       string
	description string

	help bool
	touch.Options
}

func (u *touchCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "touch", "[-acm] [-d ccyy-mm-ddTHH:MM:SS[.frac][Z]] [-r FILE] FILE ..."
	u.description = "Change file access and modification times"

	flags.BoolVar(&u.Access, "Change the access time", "-a")
	flags.BoolVar(&u.Create, "Do not create missing files", "-c")
	flags.StringVar(&u.Date, "", "Change access and modified time as RFC3339Nano", "-d")
	flags.BoolVar(&u.Modified, "Change the modified time", "-m")
	flags.StringVar(&u.Reference, "", "Use access and modified time of file", "-r")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *touchCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		err  error
		data []string
	)

	if data, err = flags.Parse(s); err != nil {
		err = fmt.Errorf("%s: %s", u.name, err)
		return err
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if err := u.Touch(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("touch", &touchCmd{})
}
