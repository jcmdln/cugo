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
// touch sets the modification and access times of files to the
// current time of day. If the file doesn't exist, it is created with
// default permissions.
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

import (
	"fmt"
	"os"
	"syscall"
	"time"

	ex "github.com/jcmdln/cugo/lib/exists"
)

var (
	Access    bool
	Create    bool
	Date      string
	Modified  bool
	Reference string

	operand    string
	finfo      os.FileInfo
	fstat      *syscall.Stat_t
	rinfo      os.FileInfo
	rstat      *syscall.Stat_t
	accessTime time.Time
	modifyTime time.Time
	date       time.Time
	err        error
)

// Touch ...
func Touch(operands []string) {
	for _, operand = range operands {
		if !Create {
			if _, err = os.Create(operand); err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			}
		}

		if ex.Exists(operand) {
			finfo, _ = os.Stat(operand)
			fstat = finfo.Sys().(*syscall.Stat_t)

			if len(Reference) > 0 {
				if rinfo, err = os.Stat(Reference); err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				}

				rstat = rinfo.Sys().(*syscall.Stat_t)
				accessTime = time.Unix(rstat.Atim.Sec, rstat.Atim.Nsec)
				modifyTime = rinfo.ModTime()
			} else if len(Date) > 0 {
				if date, err = time.Parse(time.RFC3339Nano, Date); err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				}

				accessTime = date
				modifyTime = date
			} else {
				accessTime = time.Unix(fstat.Atim.Sec, fstat.Atim.Nsec)
				modifyTime = finfo.ModTime()
			}

			if err = os.Chtimes(operand, accessTime, modifyTime); err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			}
		}
	}

	os.Exit(0)
}
