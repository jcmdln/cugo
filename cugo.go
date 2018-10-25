// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Unix utilities as a multi-call binary
//
// `cugo` provides common Unix core utilities in the form of a multi-call
// binary, with a focus on broad support for various Unix and Unix-like
// operating systems.
//
// Cugo was inspired by Rob Landley's `toybox` project with a goal of
// using Go, which allows for a much simpler build system. Go's standard
// library is feature complete enough to make many utilities trivial to
// implement, sometimes with little more than writing glue code as a
// wrapper for existing functionality.
//
// The primary source for manuals when building utilities are as follows
// - https://man.openbsd.org/
// - http://landley.net/toybox/status.html
// - https://illumos.org/man/
package main

import (
	"github.com/jcmdln/cugo/cmd"
)

func main() {
	cmd.RootCmd.Execute()
}
