// ls.go
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

//go:build testing

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/src/ls"
)

type lsCmd struct {
	ls.Options
}

func (u *lsCmd) Init() *flag.FlagSet {
	ls := flag.NewFlagSet("ls", flag.ExitOnError)
	ls.BoolVar(&u.All,
		"Show all entries other than '.' and '..'", "-A")
	ls.BoolVar(&u.Recursive, "R",
		"Recursively list directories and their entries")
	return ls
}

func (u *lsCmd) Run(s []string) error {
	if err := u.Ls(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("ls", &lsCmd{})
}
