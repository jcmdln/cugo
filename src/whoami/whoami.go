// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package whoami - return current user
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
