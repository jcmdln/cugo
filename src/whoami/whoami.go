package whoami

import (
	"fmt"
	"os/user"

	er "github.com/jcmdln/cugo/lib/error"
)

func Whoami() {
	usr, err := user.Current()
	if !er.Error("cugo", err) {
		fmt.Printf("%s\n", usr.Username)
	}
}
