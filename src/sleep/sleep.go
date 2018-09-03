package sleep

import (
	"strings"
	"time"
)

func Sleep(args string) {
	a := strings.Split(args, " ")

	for _, i := range a {
		s := strings.Split(i, "")
		if len(s) < 2 {
			i += "s"
		}

		t, _ := time.ParseDuration(i)
		time.Sleep(t)
	}
}
