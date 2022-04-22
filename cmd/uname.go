// uname.go
//
// ISC License
//
// Copyright (c) 2022 Johnathan C. Maudlin <jcmdln@gmail.com>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE
// OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

package cmd

import (
	"flag"

	"github.com/jcmdln/cugo/src/uname"
)

type unameCmd struct {
	uname.Options
}

func (u *unameCmd) Init() *flag.FlagSet {
	uname := flag.NewFlagSet("uname", flag.ExitOnError)
	uname.BoolVar(&u.All, "a", false,
		"Behave as though all options were specified")
	uname.BoolVar(&u.Machine, "m", false, "Print the machine hardware name")
	uname.BoolVar(&u.Nodename, "n", false, "Print the nodename / network name")
	uname.BoolVar(&u.Release, "r", false, "Print the operating system release")
	uname.BoolVar(&u.Sysname, "s", false, "Print the operating system name")
	uname.BoolVar(&u.Version, "v", false, "Print the operating system version")
	return uname
}

func (u *unameCmd) Run(s []string) error {
	if err := u.Uname(); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["uname"] = &unameCmd{}
}
