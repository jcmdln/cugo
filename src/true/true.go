// SPDX-License-Identifier: ISC

package true

import "os"

// True always exists with a zero exit code.
func True() {
	os.Exit(0)
}
