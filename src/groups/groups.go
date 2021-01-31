// groups.go
//
// Copyright 2021 Johnathan C. Maudlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

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
