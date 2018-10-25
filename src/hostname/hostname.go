// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// set or print name of current host system
//
// Setting the hostname is currently not supported as this is missing
// from the 'syscall' package.
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

	// name is a string used to store the hostname to be printed
	name string
)

func Hostname(args []string) {
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
