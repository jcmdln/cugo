// SPDX-License-Identifier: ISC

package hostname

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

func (opt *Options) Hostname(hostname string) error {
	var (
		err  error
		name string
		uid  int
	)

	if len(hostname) < 1 {
		if name, err = os.Hostname(); err != nil {
			return err
		}

		if opt.Strip {
			s := strings.Split(name, ".")
			name = s[0]
		}

		if _, err = fmt.Printf("%s\n", name); err != nil {
			return err
		}
	} else {
		if uid = unix.Getuid(); uid != 0 {
			err = errors.New("hostname: you must be root to change the hostname")
			return err
		}

		if err = SetHostname([]byte(hostname)); err != nil {
			return err
		}
	}

	return nil
}
