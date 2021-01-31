// doc.go
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
