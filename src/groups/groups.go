// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// show group memberships
//
// SYNOPSIS
//
//     groups [USER]
//
// DESCRIPTION
//
// groups displays the groups to which you (or the optionally specified
// user) belong.
//
// SEE ALSO
//
//     tbd
//
// REFERENCES
//
//     http://man.openbsd.org/groups
//     http://refspecs.linuxfoundation.org/LSB_4.1.0/LSB-Core-generic/LSB-Core-generic/groups.html
package groups

import (
	"fmt"
	"os"
	"os/user"
)

var (
	err error

	usr *user.User

	gid   string
	gids  []string
	gname *user.Group

	userGroups string
)

// Groups ...
func Groups(username string) {
	if len(username) == 0 {
		if usr, err = user.Current(); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		} else {
			if gids, err = usr.GroupIds(); err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			} else {
				for _, gid := range gids {
					if gname, err = user.LookupGroupId(gid); err != nil {
						fmt.Printf("cugo: %s\n", err)
						os.Exit(1)
					} else {
						userGroups += gname.Name + " "
					}
				}
			}

			fmt.Printf("%s\n", userGroups)
		}
	}

	os.Exit(0)
}
