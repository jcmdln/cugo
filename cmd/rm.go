// rm.go
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
	"github.com/jcmdln/cugo/src/rm"
	"github.com/jcmdln/flagger"
)

type rmCmd struct {
	name        string
	usage       string
	description string

	help bool
	rm.Options
}

func (u *rmCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "rm", "[-dfiPRrv] FILE ..."
	u.description = "Remove directory entries"

	flags.BoolVar(&u.Dir, "Remove empty directories", "-d")
	flags.BoolVar(&u.Force, "Skip prompts and ignore warnings", "-f")
	flags.BoolVar(&u.Interactive, "Prompt before each removal", "-i")
	flags.BoolVar(&u.Recursive, "Remove directories and their contents recursively", "-r", "-R")
	flags.BoolVar(&u.Verbose, "Print a message when actions are taken", "-v")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *rmCmd) Action(s []string, flags *flagger.Flags) error {
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

	if err = u.Rm(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("rm", &rmCmd{})
}
