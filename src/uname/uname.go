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

package uname

import (
	"fmt"
	"strings"

	"golang.org/x/sys/unix"
)

func (opt *Options) Uname() error {
	var (
		err   error
		out   []string
		uname unix.Utsname
	)

	if err = unix.Uname(&uname); err != nil {
		return err
	}

	if !opt.Nodename && !opt.Release && !opt.Version && !opt.Machine {
		opt.Sysname = true
	}

	if opt.All {
		opt.Nodename = true
		opt.Release = true
		opt.Version = true
		opt.Machine = true
	}

	if opt.Sysname {
		out = append(out, strings.Trim(string(uname.Sysname[:]), "\x00"))
	}
	if opt.Nodename {
		out = append(out, strings.Trim(string(uname.Nodename[:]), "\x00"))
	}
	if opt.Release {
		out = append(out, strings.Trim(string(uname.Release[:]), "\x00"))
	}
	if opt.Version {
		out = append(out, strings.Trim(string(uname.Version[:]), "\x00"))
	}
	if opt.Machine {
		out = append(out, strings.Trim(string(uname.Machine[:]), "\x00"))
	}

	if _, err = fmt.Printf("%s\n", strings.Join(out, " ")); err != nil {
		return err
	}

	return nil
}
