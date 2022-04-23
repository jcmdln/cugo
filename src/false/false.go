// SPDX-License-Identifier: ISC

package false

import "os"

// False always exits with a non-zero exit code.
func False() {
	os.Exit(1)
}
