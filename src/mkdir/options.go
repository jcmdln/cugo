// SPDX-License-Identifier: ISC

package mkdir

type Options struct {
	Mode    uint
	Parents bool
	Verbose bool
}

type Opts func(*Options)

func Mode(mode uint) Opts {
	return func(opt *Options) {
		opt.Mode = mode
	}
}

func Parents(unbuffered bool) Opts {
	return func(opt *Options) {
		opt.Parents = unbuffered
	}
}

func Verbose(verbose bool) Opts {
	return func(opt *Options) {
		opt.Verbose = verbose
	}
}
