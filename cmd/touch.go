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
	"flag"

	"github.com/jcmdln/cugo/src/touch"
)

type touchCmd struct {
	name        string
	usage       string
	description string

	help bool
	touch.Options
}

func (u *touchCmd) Init() *flag.FlagSet {
	touch := flag.NewFlagSet("touch", flag.ExitOnError)
	touch.BoolVar(&u.Access, "a", false, "Change the access time")
	touch.BoolVar(&u.Create, "c", false, "Do not create missing files")
	touch.StringVar(&u.Date, "d", "",
		"Change access and modified time as RFC3339Nano")
	touch.BoolVar(&u.Modified, "m", false, "Change the modified time")
	touch.StringVar(&u.Reference, "r", "",
		"Use access and modified time of file")
	return touch
}

func (u *touchCmd) Run(s []string) error {
	if err := u.Touch(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["touch"] = &touchCmd{}
}
