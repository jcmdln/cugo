// uname.go
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
