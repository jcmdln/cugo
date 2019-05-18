// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sleep

import (
	"strings"
	"time"
)

func Sleep(operands []string) error {
	var (
		dur time.Duration
		err error
		opt string
	)

	for _, opt = range operands {
		if len(strings.Split(opt, "")) < 2 {
			opt += "s"
		}

		if dur, err = time.ParseDuration(opt); err != nil {
			return err
		}

		time.Sleep(dur)
	}

	return nil
}
