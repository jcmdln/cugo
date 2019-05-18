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
		err   error
		gid   string
		gids  []string
		group *user.Group
		out   string
		usr   *user.User
	)

	if len(username) < 1 {
		if usr, err = user.Current(); err != nil {
			return err
		}
	} else {
		if usr, err = user.Lookup(username); err != nil {
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

	if _, err = fmt.Printf("%s\n", out); err != nil {
		return err
	}

	return nil
}
