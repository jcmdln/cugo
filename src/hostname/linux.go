// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

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
