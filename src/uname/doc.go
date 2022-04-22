// doc.go
//
// ISC License
//
// Copyright (c) 2022 Johnathan C. Maudlin <jcmdln@gmail.com>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE
// OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

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
