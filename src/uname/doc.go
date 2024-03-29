// SPDX-License-Identifier: ISC

// print operating system information.
//
// SYNOPSIS
//
//     uname [-amnprsv]
//
// DESCRIPTION
//
// uname writes symbols representing one or more system characteristics
// to stdout.
//
// The options are as follows:
//
//     -a        Behave as though all options were specified
//
//     -m        Print the machine hardware name
//
//     -n        Print the nodename (aka network name)
//
//     -r        Print the operating system release
//
//     -s        Print the operating system name
//
//     -v        Print the operating system version
//
// If no options are specified, uname prints the operating system name
// as if '-s' had been specified.
//
// SEE ALSO
//
//     WIP
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/uname.html
//
package uname
