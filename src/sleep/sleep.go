// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// delay for a specified amount of time
//
// SYNOPSIS
//
//    sleep DURATION ...
//
// DESCRIPTION
//
// Sleep suspends execution for a minimum of the specified duration of
// time. By default, sleep uses seconds as it's unit of time though
// supports seconds (s), minutes (m), and hours(m) as characters at the
// end of the provided interval. This number must be positive and may
// contain a decimal.
//
// SEE ALSO
//
// * https://golang.org/pkg/time/#Sleep
//
// REFERENCES
//
// * http://man.openbsd.org/sleep
// * http://pubs.opengroup.org/onlinepubs/9699919799/utilities/sleep.html
package sleep

import (
	"fmt"
	"os"
	"strings"
	"time"
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
		if err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		} else {
			time.Sleep(t)
		}
	}
}
