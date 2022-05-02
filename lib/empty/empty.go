// SPDX-License-Identifier: ISC

package empty

import (
	"io"
	"os"
)

// Empty confirms whether the provided directory contains any children,
// returning 'true' if a child file or directory is present.
func Empty(dir string) bool {
	var (
		err  error
		file *os.File
	)

	if file, err = os.Open(dir); err != nil {
		return true
	}
	defer file.Close()

	if _, err = file.Readdirnames(1); err == io.EOF {
		return true
	}

	return false
}
