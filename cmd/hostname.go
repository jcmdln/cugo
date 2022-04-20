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
	"flag"

	"github.com/jcmdln/cugo/src/hostname"
)

type hostnameCmd struct {
	hostname.Options
}

func (u *hostnameCmd) Init() *flag.FlagSet {
	hostname := flag.NewFlagSet("hostname", flag.ExitOnError)
	hostname.BoolVar(&u.Strip, "s", false,
		"Trim any domain information from the printed name")
	return hostname
}

func (u *hostnameCmd) Run(s []string) error {
	var hostname string

	if len(s) == 0 {
		hostname = ""
	} else {
		hostname = s[0]
	}

	if err := u.Hostname(hostname); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["hostname"] = &hostnameCmd{}
}
