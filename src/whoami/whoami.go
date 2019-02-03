// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// return current user
//
// SYNOPSIS
//
//     whoami
//
// DESCRIPTION
//
// whoami displays your effective user ID as a name.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/user/#User
//     https://golang.org/pkg/os/user/#Current
//
// REFERENCES
//
//     http://man.openbsd.org/whoami
//     https://linux.die.net/man/1/whoami
package whoami

import (
	"fmt"
	"os"
	"os/user"
)

// Whoami displays your effective user ID as a name.
func Whoami() {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("cugo: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", usr.Username)
	}
}
