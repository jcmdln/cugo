// hostname.go
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

package hostname

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

func (opt *Options) Hostname(hostname string) error {
	var (
		err  error
		name string
		uid  int
	)

	if len(hostname) < 1 {
		if name, err = os.Hostname(); err != nil {
			return err
		}

		if opt.Strip {
			s := strings.Split(name, ".")
			name = s[0]
		}

		if _, err = fmt.Printf("%s\n", name); err != nil {
			return err
		}
	} else {
		if uid = unix.Getuid(); uid != 0 {
			err = errors.New("hostname: you must be root to change the hostname")
			return err
		}

		if err = SetHostname([]byte(hostname)); err != nil {
			return err
		}
	}

	return nil
}
