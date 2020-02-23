// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

// +build openbsd

package hostname

import (
	"errors"
)

func SetHostname(hostname []byte) error {
	err := errors.New("hostname: setting the hostname is not available on this platform")
	return err
}
