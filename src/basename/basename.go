// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// return non-directory portion of a pathname
//
// SYNOPSIS
//
//     basename string [suffix]
//
// DESCRIPTION
//
// basename deletes any prefix ending with the last slash (‘/’)
// character present in string, and a suffix, if given. The resulting
// filename is written to the standard output. A non-existent suffix is
// ignored.
//
// SEE ALSO
//
//     https://golang.org/pkg/path/filepath/#Base
//
// REFERENCES
//
//     http://man.openbsd.org/basename
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/basename.html
package basename

import (
	"fmt"
	"os"
	"path/filepath"
)

func Basename(target string) {
	fmt.Printf("%s\n", filepath.Base(target))

	os.Exit(0)
}
