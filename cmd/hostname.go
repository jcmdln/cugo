// hostname.go
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
