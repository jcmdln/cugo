// prompt.go
//
// Copyright 2021 Johnathan C. Maudlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

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
