// SPDX-License-Identifier: ISC

//go:build testing
// +build testing

package tail

type Options struct {
	Count  int
	Follow bool
	Number int
}

type Opts func(*Options)
