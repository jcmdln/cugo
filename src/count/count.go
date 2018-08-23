package count

import "fmt"

func Count(args []string) {
	var total int
	for i, _ := range args {
		total = i + 1
	}
	fmt.Printf("%v\n", total)
}
