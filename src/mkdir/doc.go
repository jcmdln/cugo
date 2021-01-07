// Copyright 2021 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

// create directories.
//
// SYNOPSIS
//
//     mkdir [-p] [-m MODE] DIRECTORY ...
//
// DESCRIPTION
//
// mkdir creates the directories named as arguments in the order
// specified using mode rwxrwxrwx (0777) as modified by the current
// umask.
//
// The options are as follows:
//
//     -m        Set the file permission bits of a new directory to MODE.
//
//     -p        Create missing parent directories as required.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#FileMode
//     https://golang.org/pkg/os/#Mkdir
//     https://golang.org/pkg/os/#MkdirAll
//
// REFERENCES
//
//     https://pubs.opengroup.org/onlinepubs/9699919799/utilities/mkdir.html
//
package mkdir
