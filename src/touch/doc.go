// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// change file access and modification times
//
// SYNOPSIS
//
//     touch [-acm] [-d ccyy-mm-ddTHH:MM:SS[.frac][Z]] [-r FILE] FILE ...
//
// DESCRIPTION
//
// touch sets the modification and access times of files to the current
// time of day. If the file doesn't exist, it is created with default
// permissions.
//
// The options are as follows:
//
//     -a        Change the access time of the file.
//
//     -c        Do not create missing files, or display an error when a
//               file is either missing or not created.
//
//     -d        Change the access and modified times of the file, using
//               ISO8601/RFC3339Nano format.
//
//     -m        Change the modified time of the file.
//
//     -r        Use the access and modified times of the reference file
//               rather than the current date-time.
//
// SEE ALSO
//
//     https://golang.org/pkg/time/#pkg-constants
//     https://golang.org/pkg/os/#Chtimes
//     https://golang.org/pkg/time/#Unix
//
// REFERENCES
//
//     http://man.openbsd.org/touch
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/touch.html
package touch
