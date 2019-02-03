// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// return false value
//
// SYNOPSIS
//
//     false
//
// DESCRIPTION
//
// false always exits with a non-zero exit code.
//
// SEE ALSO
//
// * https://golang.org/pkg/os/#Exit
//
// REFERENCES
//
// * http://man.openbsd.org/false
// * http://pubs.opengroup.org/onlinepubs/9699919799/utilities/false.html
package false

import "os"

// False always exits with a non-zero exit code.
func False() {
	os.Exit(1)
}
