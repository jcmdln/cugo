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
// * https://golang.org/pkg/time/#pkg-constants
// * https://golang.org/pkg/os/#Chtimes
// * https://golang.org/pkg/time/#Unix
//
// REFERENCES
//
// * http://man.openbsd.org/touch
// * http://pubs.opengroup.org/onlinepubs/9699919799/utilities/touch.html
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
)

func touch(file string, atime time.Time, mtime time.Time) {
	f, _ := os.Stat(file)
	fstat := f.Sys().(*syscall.Stat_t)

	if Access && Modified || !Access && !Modified {
		err := os.Chtimes(file, atime, mtime)
		if err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}
	} else {
		if Access {
			err := os.Chtimes(file, atime, f.ModTime())
			if err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			}
		}

		if Modified {
			err := os.Chtimes(file, time.Unix(fstat.Atim.Sec, fstat.Atim.Nsec), mtime)
			if err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			}
		}
	}
}

func Touch(args []string) {
	for _, file := range args {
		if Create {
			_, err := os.Create(file)
			if err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			}
		}

		if ex.Exists(file) {
			if len(Reference) > 0 {
				r, err := os.Stat(Reference)
				if err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				} else {
					rstat := r.Sys().(*syscall.Stat_t)
					touch(file, time.Unix(rstat.Atim.Sec, rstat.Atim.Nsec), r.ModTime())
				}
			} else if len(Date) > 0 {
				date, err := time.Parse(time.RFC3339Nano, Date)
				if err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				} else {
					touch(file, date, date)
				}
			} else {
				f, err := os.Stat(file)
				if err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				} else {
					fstat := f.Sys().(*syscall.Stat_t)
					touch(file, time.Unix(fstat.Atim.Sec, fstat.Atim.Nsec), f.ModTime())
				}
			}
		}
	}
}
