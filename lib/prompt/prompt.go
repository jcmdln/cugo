package prompt

import "fmt"

func Prompt(text string) bool {
	fmt.Printf(text + " [Yes/No]: ")
	var a string

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if a == "y" || a == "Y" || a == "yes" || a == "Yes" {
		return true
	}

	return false
}
