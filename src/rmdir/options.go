// SPDX-License-Identifier: ISC

package rmdir

type Options struct {
	Parents bool
	Verbose bool
}

type Opts func(*Options)

func Parents(parents bool) Opts {
	return func(opt *Options) {
		opt.Parents = parents
	}
}

func Verbose(verbose bool) Opts {
	return func(opt *Options) {
		opt.Verbose = verbose
	}
}
