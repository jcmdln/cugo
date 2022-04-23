// SPDX-License-Identifier: ISC

package prompt

import "fmt"

// Prompt provides a simple "yes|no" prompt to the user and checks their
// response. If the user answers the prompt with anything other than
// Y|y|Yes|yes, return 'false'.
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
