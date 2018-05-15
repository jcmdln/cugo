package sleep

import (
	"strings"
	"time"
)

func Sleep(args []string) {
	for _, i := range args {
		s := strings.Split(i, "")
		if len(s) < 2 {
			i += "s"
		}
		t, _ := time.ParseDuration(i)
		time.Sleep(t)
	}
}
