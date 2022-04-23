// SPDX-License-Identifier: ISC

package touch

import (
	"os"
	"syscall"
	"time"

	ex "github.com/jcmdln/cugo/lib/exists"
)

type Option struct {
	Access    bool
	Create    bool
	Date      string
	Modified  bool
	Reference string
}

// Touch ...
func (opt *Option) Touch(operands []string) error {
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
				return err
			}
		}

		if err = ex.Exists(operand); err != nil {
			finfo, _ = os.Stat(operand)
			fstat = finfo.Sys().(*syscall.Stat_t)

			if len(opt.Reference) > 0 {
				if rinfo, err = os.Stat(opt.Reference); err != nil {
					return err
				}

				rstat = rinfo.Sys().(*syscall.Stat_t)
				accessTime = time.Unix(rstat.Atim.Sec, rstat.Atim.Nsec)
				modifyTime = rinfo.ModTime()
			} else if len(opt.Date) > 0 {
				if date, err = time.Parse(time.RFC3339Nano, opt.Date); err != nil {
					return err
				}

				accessTime = date
				modifyTime = date
			} else {
				accessTime = time.Unix(fstat.Atim.Sec, fstat.Atim.Nsec)
				modifyTime = finfo.ModTime()
			}

			if err = os.Chtimes(operand, accessTime, modifyTime); err != nil {
				return err
			}
		}
	}

	return nil
}
