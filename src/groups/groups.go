// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package groups

import (
	"fmt"
	"os"
	"os/user"
)

var (
	err       error
	gid       string
	gids      []string
	gname     *user.Group
	usr       *user.User
	usrGroups string
)

// Groups ...
func Groups(username string) {
	if len(username) == 0 {
		if usr, err = user.Current(); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}

		if gids, err = usr.GroupIds(); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}

		for _, gid = range gids {
			if gname, err = user.LookupGroupId(gid); err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			}

			usrGroups += gname.Name + " "
		}

		fmt.Printf("%s\n", usrGroups)
	}

	os.Exit(0)
}
