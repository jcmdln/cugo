// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

// +build linux

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
