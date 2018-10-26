// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// set or print name of current host system
//
// Setting the hostname is currently not supported as this is missing
// from the 'syscall' package.
//
// hostname is used to set or print the name of the current host. If no
// argument is given, the name of the current host is printed.
//
// The host name may be set by the superuser by specifying a hostname or
// supplying a myname or hostname file, which is used at boot.
//
// Available options:
//
//     -s, strip    trims domain information from the printed name.
//
package hostname

import (
	"fmt"
	"os"
	"strings"

	er "github.com/jcmdln/cugo/lib/error"
)

var (
	// Strip is a bool that when true removes domain information
	Strip bool
)

func Hostname(args []string) {
	var name string

	name, err := os.Hostname()
	er.Error("cugo", err)

	if Strip {
		s := strings.Split(name, ".")
		name = s[0]
	}

	if len(args) < 1 {
		fmt.Printf("%s\n", name)
	}
}
