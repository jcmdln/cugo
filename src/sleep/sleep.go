// SPDX-License-Identifier: ISC

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
