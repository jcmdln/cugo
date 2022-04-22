// mkdir.go
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

	"github.com/jcmdln/cugo/src/mkdir"
)

type mkdirCmd struct {
	mkdir.Options
}

func (u *mkdirCmd) Init() *flag.FlagSet {
	mkdir := flag.NewFlagSet("mkdir", flag.ExitOnError)
	mkdir.UintVar(&u.Mode, "m", 0755, "Set permissions to MODE value")
	mkdir.BoolVar(&u.Parents, "p", false, "Create missing parent directories")
	mkdir.BoolVar(&u.Verbose, "v", false, "Display each created directory")
	return mkdir
}

func (u *mkdirCmd) Run(s []string) error {
	if _, err := u.Mkdir(s); err != nil {
		return err
	}

	return nil
}

func init() {
	Commands["mkdir"] = &mkdirCmd{}
}
