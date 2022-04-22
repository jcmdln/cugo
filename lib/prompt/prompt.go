// prompt.go
//
// ISC License
//
// Copyright (c) 2022 Johnathan C. Maudlin <jcmdln@gmail.com>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE
// OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

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
