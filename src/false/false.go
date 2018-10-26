// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// return false value
//
// false always exits with a non-zero exit code.
package false

import "os"

func False() {
	os.Exit(1)
}
