// Copyright 2020 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that may
// be found in the LICENSE file.

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
