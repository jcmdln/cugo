package touch

import (
	"fmt"
	"os"

	e "github.com/jcmdln/cugo/lib/exists"
)

var (
	Access    bool
	Create    bool
	Modified  bool
	Reference string
	Time      string
	Verbose   bool
)

func Touch(args []string) {
	for _, file := range args {
		if !e.Exists(file) {
			if !Create {
				os.Create(file)
				if Verbose {
					fmt.Printf("cugo: touch: Created '%s'\n", file)
				}
			}
		}
	}
}
