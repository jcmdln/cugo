// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

	if !opt.All && !opt.Nodename && !opt.Release && !opt.Version && !opt.Machine {
		opt.Sysname = true
	} else {
		if opt.All {
			opt.Sysname = true
			opt.Nodename = true
			opt.Release = true
			//opt.Version = true
			opt.Machine = true
		}
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
