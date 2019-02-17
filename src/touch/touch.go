// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package touch

import (
	"fmt"
	"os"
	"syscall"
	"time"

	ex "github.com/jcmdln/cugo/lib/exists"
)

type Options struct {
	Access    bool
	Create    bool
	Date      string
	Modified  bool
	Reference string
}

type Opts func(*Options)

func Access(access bool) Opts {
	return func(opt *Options) {
		opt.Access = access
	}
}

func Create(create bool) Opts {
	return func(opt *Options) {
		opt.Create = create
	}
}

func Date(Date string) Opts {
	return func(opt *Options) {
		opt.Date = Date
	}
}

func Modified(modified bool) Opts {
	return func(opt *Options) {
		opt.Modified = modified
	}
}

func Reference(reference string) Opts {
	return func(opt *Options) {
		opt.Reference = reference
	}
}

// Touch ...
func (opt *Options) Touch(operands []string) {
	var (
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

	for _, operand = range operands {
		if !opt.Create {
			if _, err = os.Create(operand); err != nil {
				fmt.Printf("cugo: %s\n", err)
				os.Exit(1)
			}
		}

		if ex.Exists(operand) {
			finfo, _ = os.Stat(operand)
			fstat = finfo.Sys().(*syscall.Stat_t)

			if len(opt.Reference) > 0 {
				if rinfo, err = os.Stat(opt.Reference); err != nil {
					fmt.Printf("cugo: %s\n", err)
					os.Exit(1)
				}

				rstat = rinfo.Sys().(*syscall.Stat_t)
				accessTime = time.Unix(rstat.Atim.Sec, rstat.Atim.Nsec)
				modifyTime = rinfo.ModTime()
			} else if len(opt.Date) > 0 {
				if date, err = time.Parse(time.RFC3339Nano, opt.Date); err != nil {
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
