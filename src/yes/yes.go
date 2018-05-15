package yes

import (
	"io"
	"os"
	"strings"
)

func Yes(args []string) {
	if len(args) == 0 {
		for true {
			io.WriteString(os.Stdout, "y\n")
		}
	} else {
		for true {
			out := strings.Join(args, " ") + "\n"
			io.WriteString(os.Stdout, out)
		}
	}
}
