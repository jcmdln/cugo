// SPDX-License-Identifier: ISC

package whoami

import (
	"fmt"
	"os"
	"os/user"
)

// Whoami displays your effective user ID as a name.
func Whoami() {
	if usr, err := user.Current(); err != nil {
		fmt.Printf("cugo: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", usr.Username)
	}

	os.Exit(0)
}
