// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// delay for a specified amount of time
//
// sleep suspends execution for a minimum of the specified duration of
// time. By default, sleep uses seconds as it's unit of time though
// supports seconds (s), minutes (m), and hours(m) as characters at the
// end of the provided interval. This number must be positive and may
// contain a decimal.
//
// Examples:
//
//     1, 1s    sleep for one second
//     1m       sleep for one minute
//     1h       sleep for one hour
//
package sleep

import (
	"strings"
	"time"

	er "github.com/jcmdln/cugo/lib/error"
)

func Sleep(args string) {
	a := strings.Split(args, " ")

	var (
		s   []string
		t   time.Duration
		err error
	)

	for _, i := range a {
		s = strings.Split(i, "")
		if len(s) < 2 {
			i += "s"
		}

		t, err = time.ParseDuration(i)
		if !er.Error("cugo", err) {
			time.Sleep(t)
		}
	}
}
