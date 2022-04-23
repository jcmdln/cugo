// SPDX-License-Identifier: ISC

//go:build testing
// +build testing

package ls

type Options struct {
	All       bool
	Recursive bool
}

type Opts func(*Options)

func All(all bool) Opts {
	return func(opt *Options) {
		opt.All = all
	}
}

func Recursive(recursive bool) Opts {
	return func(opt *Options) {
		opt.Recursive = recursive
	}
}
