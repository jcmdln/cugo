// SPDX-License-Identifier: ISC

// set or print the system hostname.
//
// SYNOPSIS
//
//     hostname [-s] [HOSTNAME]
//
// DESCRIPTION
//
// hostname utility is used to set or print the name of the current
// host.  If no argument is given, the name of the current host is
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
//     https://refspecs.linuxfoundation.org/LSB_4.1.0/LSB-Core-generic/LSB-Core-generic/hostname.html
//     https://man.openbsd.org/hostname.1
//
package hostname
