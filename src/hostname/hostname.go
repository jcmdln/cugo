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

type Options struct {
	Strip bool
}

type Opts func(*Options)

func Strip(strip bool) Opts {
	return func(opt *Options) {
		opt.Strip = strip
	}
}
