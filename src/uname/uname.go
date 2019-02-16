// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// print operating system information
//
// SYNOPSIS
//
//     uname [-amnprsv]
//
// DESCRIPTION
//
// uname writes symbols representing one or more system characteristics
// to stdout.
//
// The options are as follows:
//
//     -a        Behave as though all options were specified
//
//     -m        Print the machine hardware name
//
//     -n        Print the nodename (aka network name)
//
//     -r        Print the operating system release
//
//     -s        Print the operating system name
//
//     -v        Print the operating system version
//
// If no options are specified, uname prints the operating system name
// as if '-s' had been specified.
//
// SEE ALSO
//
//     WIP
//
// REFERENCES
//
//     http://man.openbsd.org/uname
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/uname.html
package uname

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

type Options struct {
	All      bool // -a
	Machine  bool // -m
	Nodename bool // -n
	Release  bool // -r
	Sysname  bool // -s
	Version  bool // -v
}

type Opts func(*Options)

func All(all bool) Opts {
	return func(opt *Options) {
		opt.All = all
	}
}

func Machine(machine bool) Opts {
	return func(opt *Options) {
		opt.Machine = machine
	}
}

func Nodename(nodename bool) Opts {
	return func(opt *Options) {
		opt.Nodename = nodename
	}
}

func Release(release bool) Opts {
	return func(opt *Options) {
		opt.Release = release
	}
}

func Sysname(sysname bool) Opts {
	return func(opt *Options) {
		opt.Sysname = sysname
	}
}

func Version(version bool) Opts {
	return func(opt *Options) {
		opt.Version = version
	}
}

func (opt *Options) Uname() {
	var (
		uname unix.Utsname
		list  []string
	)

	if err := unix.Uname(&uname); err != nil {
		fmt.Printf("cugo: uname: %s\n", err)
	}

	if opt.All {
		fmt.Printf("%s %s %s %s %s\n", uname.Sysname, uname.Nodename,
			uname.Release, uname.Version, uname.Machine)
		os.Exit(0)
	}

	if !opt.All && !opt.Sysname && !opt.Nodename &&
		!opt.Release && !opt.Version && !opt.Machine {
		fmt.Printf("%s\n", uname.Sysname)
		os.Exit(0)
	}

	if opt.Sysname {
		list = append(list, string(uname.Sysname[:]))
	}
	if opt.Nodename {
		list = append(list, string(uname.Nodename[:]))
	}
	if opt.Release {
		list = append(list, string(uname.Release[:]))
	}
	if opt.Version {
		list = append(list, string(uname.Version[:]))
	}
	if opt.Machine {
		list = append(list, string(uname.Machine[:]))
	}

	fmt.Printf("%s\n", strings.Join(list, " "))

	os.Exit(0)
}
