// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package false - return false value
package false

import "os"

// False always exits with a non-zero exit code.
func False() {
	os.Exit(1)
}
