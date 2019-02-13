// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

// +build linux

package hostname

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

func (opt *Options) Hostname(hostname string) {
	var (
		name string
		err  error
	)

	if len(hostname) == 0 {
		if name, err = os.Hostname(); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}

		if opt.Strip {
			s := strings.Split(name, ".")
			name = s[0]
		}

		fmt.Printf("%s\n", name)
	} else {
		if uid := unix.Getuid(); uid != 0 {
			fmt.Println("cugo: hostname: you must be root to change the hostname")
		} else {
			unix.Sethostname([]byte(hostname))
		}
	}

	os.Exit(0)
}
