// SPDX-License-Identifier: ISC

//go:build linux

package hostname

import (
	"golang.org/x/sys/unix"
)

func SetHostname(hostname []byte) error {
	if err := unix.Sethostname([]byte(hostname)); err != nil {
		return err
	}

	return nil
}
