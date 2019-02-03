// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// set or print the system hostname
//
// SYNOPSIS
//
//     hostname [-s] [HOSTNAME]
//
// DESCRIPTION
//
// hostname utility is used to set or print the name of the current
// host. If no argument is given, the name of the current host is
// printed.
//
// The options are as follows:
//
//     -s        Trim any domain information from the printed name.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#Hostname
//
// REFERENCES
//
//     http://man.openbsd.org/hostname
//     https://linux.die.net/man/1/hostname
package hostname

import (
	"fmt"
	"os"
	"strings"
)

var (
	// Strip is a bool that when true removes domain information
	Strip bool
)

// Hostname is used to set or print the system hostname. If no argument
// is given, the current system hostname is printed.
//
// The hostname currently can not be set, as it has not yet been
// implemented.
func Hostname(hostname string) {
	var name string

	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("cugo: %s\n", err)
		os.Exit(1)
	}

	if Strip {
		s := strings.Split(name, ".")
		name = s[0]
	}

	fmt.Printf("%s\n", name)
}
