package touch

import (
	"os"
	"syscall"
	"time"

	er "github.com/jcmdln/cugo/lib/error"
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
		er.Error("cugo", err)
	} else {
		if Access {
			err := os.Chtimes(file, atime, f.ModTime())
			er.Error("cugo", err)
		}

		if Modified {
			err := os.Chtimes(file, time.Unix(fstat.Atim.Sec, fstat.Atim.Nsec), mtime)
			er.Error("cugo", err)
		}
	}
}

func Touch(args []string) {
	for _, file := range args {
		if Create {
			_, err := os.Create(file)
			er.Error("cugo", err)
		}

		if ex.Exists(file) {
			if len(Reference) > 0 {
				r, err := os.Stat(Reference)
				if !er.Error("cugo", err) {
					rstat := r.Sys().(*syscall.Stat_t)
					touch(file, time.Unix(rstat.Atim.Sec, rstat.Atim.Nsec), r.ModTime())
				}
			} else if len(Date) > 0 {
				date, err := time.Parse(time.RFC3339Nano, Date)
				if !er.Error("cugo", err) {
					touch(file, date, date)
				}
			} else {
				f, err := os.Stat(file)
				if !er.Error("cugo", err) {
					fstat := f.Sys().(*syscall.Stat_t)
					touch(file, time.Unix(fstat.Atim.Sec, fstat.Atim.Nsec), f.ModTime())
				}
			}
		}
	}
}
