// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package uname

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

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
