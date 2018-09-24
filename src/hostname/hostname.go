package hostname

import (
	"fmt"
	"os"

	er "github.com/jcmdln/cugo/lib/error"
)

func Hostname(args []string) {
	h, err := os.Hostname()
	if !er.Error("cugo", err) {
		fmt.Printf("%s\n", h)
	}
}
