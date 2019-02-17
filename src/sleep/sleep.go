// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sleep

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Sleep(operands []string) {
	var (
		opt string
		dur time.Duration
		err error
	)

	for _, opt = range operands {
		if len(strings.Split(opt, "")) < 2 {
			opt += "s"
		}

		if dur, err = time.ParseDuration(opt); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		}

		time.Sleep(dur)
	}

	os.Exit(0)
}
