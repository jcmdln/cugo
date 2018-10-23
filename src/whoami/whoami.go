// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package whoami

import (
	"fmt"
	"os/user"

	er "github.com/jcmdln/cugo/lib/error"
)

func Whoami() {
	usr, err := user.Current()
	if !er.Error("cugo", err) {
		fmt.Printf("%s\n", usr.Username)
	}
}
