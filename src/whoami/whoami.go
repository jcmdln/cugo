package whoami

import (
	"fmt"
	"os/user"
)

func Whoami() {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("cugo: %s", err)
		return
	}
	fmt.Printf("%s\n", usr.Username)
}
