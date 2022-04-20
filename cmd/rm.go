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
	"flag"

	"github.com/jcmdln/cugo/src/rm"
)

type rmCmd struct {
	rm.Options
}

func (u *rmCmd) Init() *flag.FlagSet {
	rm := flag.NewFlagSet("rm", flag.ExitOnError)
	rm.BoolVar(&u.Dir, "d", false, "Remove empty directories")
	rm.BoolVar(&u.Force, "f", false, "Skip prompts and ignore warnings")
	rm.BoolVar(&u.Interactive, "i", false, "Prompt before each removal")
	rm.BoolVar(&u.Recursive, "r", false,
		"Remove directories and their contents recursively")
	rm.BoolVar(&u.Verbose, "v", false,
		"Print a message when actions are taken")
	return rm
}

func (u *rmCmd) Run(s []string) error {
	if err := u.Rm(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["rm"] = &rmCmd{}
}
