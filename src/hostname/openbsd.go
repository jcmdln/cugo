// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

// +build openbsd

package hostname

import (
	"fmt"
	"os"
	"strings"
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
		// Setting the hostname is not implemented in OpenBSD
	}

	os.Exit(0)
}
