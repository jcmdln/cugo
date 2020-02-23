// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package false

import "os"

// False always exits with a non-zero exit code.
func False() {
	os.Exit(1)
}
