// rm.go
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
