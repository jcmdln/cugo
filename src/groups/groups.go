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
		gid   string
		gids  []string
		group *user.Group
		usr   *user.User

		err error
		out string
	)

	if len(username) < 1 {
		usr, err = user.Current()
		if err != nil {
			return err
		}
	} else {
		usr, err = user.Lookup(username)
		if err != nil {
			return err
		}

		out = username + " : "
	}

	if gids, err = usr.GroupIds(); err != nil {
		return err
	}

	for _, gid = range gids {
		if group, err = user.LookupGroupId(gid); err != nil {
			return err
		}

		out += group.Name + " "
	}

	fmt.Printf("%s\n", out)

	return nil
}
