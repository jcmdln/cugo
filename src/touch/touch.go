package touch

import (
	"fmt"
	"os"
	"syscall"
	"time"

	e "github.com/jcmdln/cugo/lib/exists"
)

var (
	Access    bool
	Create    bool
	Date      string
	Modified  bool
	Reference string
)

func touch(file string, atime time.Time, mtime time.Time) {
	f, _ := os.Stat(Reference)
	fstat := f.Sys().(*syscall.Stat_t)
	catime := time.Unix(fstat.Atim.Sec, fstat.Atim.Nsec)
	cmtime := f.ModTime()

	if Access && Modified {
		os.Chtimes(file, atime, mtime)
	} else {
		if Access {
			os.Chtimes(file, atime, cmtime)
		}

		if Modified {
			os.Chtimes(file, catime, mtime)
		}
	}
}

func Touch(args []string) {
	for _, file := range args {
		if Create {
			_, err := os.Create(file)
			if err != nil {
				fmt.Printf("cugo: %s\n", err)
			}
		}

		if e.Exists(file) {
			if len(Reference) > 0 {
				r, err := os.Stat(Reference)
				if err != nil {
					fmt.Printf("cugo: touch %s: no such file or directory\n", Reference)
					return
				}

				rstat := r.Sys().(*syscall.Stat_t)
				ratime := time.Unix(rstat.Atim.Sec, rstat.Atim.Nsec)
				rmtime := r.ModTime()
				touch(file, ratime, rmtime)
			} else {
				// No reference file
				// f, _ := os.Stat(file)
				// fstat := f.Sys().(*syscall.Stat_t)
				// fatime := time.Unix(fstat.Atim.Sec, fstat.Atim.Nsec)
				// fmtime := f.ModTime()
				// touch(file, fatime, fmtime)
			}
		}
	}
}
