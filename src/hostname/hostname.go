// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

package hostname

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

func (opt *Options) Hostname(hostname string) error {
	if len(hostname) < 1 {
		name, err := os.Hostname()
		if err != nil {
			return err
		}

		if opt.Strip {
			s := strings.Split(name, ".")
			name = s[0]
		}

		fmt.Printf("%s\n", name)
	} else {
		uid := unix.Getuid()
		if uid != 0 {
			err := errors.New("hostname: you must be root to change the hostname")
			return err
		}

		if err := SetHostname([]byte(hostname)); err != nil {
			return err
		}
	}

	return nil
}
