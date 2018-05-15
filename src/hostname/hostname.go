package hostname

import (
	"fmt"
	"os"
)

func Hostname(args []string) {
	h, _ := os.Hostname()
	fmt.Printf("%s\n", h)
}
