// whoami.go
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
	"strings"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/whoami"
	"github.com/jcmdln/flagger"
)

type whoamiCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *whoamiCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "whoami", ""
	u.description = "Display effective user ID"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *whoamiCmd) Action(s []string, flags *flagger.Flags) error {
	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if len(s) > 0 {
		t := strings.Join(s, " ")
		fmt.Printf("cugo: whoami: extra operand '%s'\n", t)
	} else {
		whoami.Whoami()
	}

	return nil
}

func init() {
	Command.Add("whoami", &whoamiCmd{})
}
