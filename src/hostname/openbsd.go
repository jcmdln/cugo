// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

// +build openbsd

package hostname

import (
	"errors"
)

func SetHostname(hostname []byte) error {
	err := errors.New("hostname: setting the hostname is not available on this platform")
	return err
}
