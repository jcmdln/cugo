// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package groups

import (
	"fmt"
	"os/user"
)

// Groups ...
func Groups(username string) error {
	var (
		gid       string
		gids      []string
		gname     *user.Group
		usr       *user.User
		usrGroups string

		err error
	)

	if len(username) == 0 {
		if usr, err = user.Current(); err != nil {
			return err
		}

		if gids, err = usr.GroupIds(); err != nil {
			return err
		}

		for _, gid = range gids {
			if gname, err = user.LookupGroupId(gid); err != nil {
				return err
			}

			usrGroups += gname.Name + " "
		}

		fmt.Printf("%s\n", usrGroups)
	}

	return nil
}
