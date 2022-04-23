// SPDX-License-Identifier: ISC

package empty

import (
	"io"
	"os"
)

// Empty confirms whether the provided directory contains any children,
// returning 'true' if a child file or directory is present.
func Empty(dir string) bool {
	t, err := os.Open(dir)
	defer t.Close()

	_, err = t.Readdirnames(1)
	if err == io.EOF {
		return true
	}

	return false
}
