// rmdir.go
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

	"github.com/jcmdln/cugo/src/rmdir"
)

type rmdirCmd struct {
	rmdir.Options
}

func (u *rmdirCmd) Init() *flag.FlagSet {
	rmdir := flag.NewFlagSet("rmdir", flag.ExitOnError)
	rmdir.BoolVar(&u.Parents, "p", false, "Remove parent directories")
	rmdir.BoolVar(&u.Verbose, "v", false,
		"Print a message when actions are taken")
	return rmdir
}

func (u *rmdirCmd) Run(s []string) error {
	if err := u.Rmdir(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["rmdir"] = &rmdirCmd{}
}
