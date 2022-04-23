// SPDX-License-Identifier: ISC

//go:build openbsd

package hostname

import (
	"errors"
)

func SetHostname(hostname []byte) error {
	err := errors.New("hostname: setting the hostname is not available on this platform")
	return err
}
