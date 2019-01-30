// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// change file access and modification times
//
// touch sets the modification and access times of files to the current
// time. If the specified file doesn't exists, it is created with default
// permissions.
//
// Available options:
//
//     -a, --access       Change the access time
//
//     -c, --create       Do not create missing files
//
//     -d, --date         Change access and modified time as per ISO8601
//
//     -m, --modified     Change the modified time
//
//     -r, --reference    Use access/modified times from reference file
//
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
