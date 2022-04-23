// SPDX-License-Identifier: ISC

// return non-directory portion of a pathname.
//
// SYNOPSIS
//
//     basename STRING [suffix]
//
// DESCRIPTION
//
// basename deletes any prefix ending with the last slash ("/")
// character present in STRING, and a suffix, if given.  The resulting
// filename is written to the standard output.  A non-existent suffix
// is ignored.
//
// SEE ALSO
//
//     https://golang.org/pkg/path/filepath/#Base
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/basename.html
//
package basename

import (
	"path/filepath"
	"strings"
)

// Basename returns the non-directory portion of a pathname by deleting
// any prefix ending with the last forward slash (‘/’), and a suffix if
// given which may be an empty string if not required.
func Basename(path string, suffix string) (string, error) {
	var s string

	s = filepath.Base(path)
	if len(suffix) > 0 {
		s = strings.TrimSuffix(s, suffix)
	}

	return s, nil
}
