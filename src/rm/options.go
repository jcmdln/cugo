// SPDX-License-Identifier: ISC

package rm

type Options struct {
	Dir         bool
	Force       bool
	Interactive bool
	Recursive   bool
	Verbose     bool
}

type Opts func(*Options)

func Dir(dir bool) Opts {
	return func(opt *Options) {
		opt.Dir = dir
	}
}

func Force(force bool) Opts {
	return func(opt *Options) {
		opt.Force = force
	}
}

func Interactive(interactive bool) Opts {
	return func(opt *Options) {
		opt.Interactive = interactive
	}
}

func Recursive(recursive bool) Opts {
	return func(opt *Options) {
		opt.Recursive = recursive
	}
}

func Verbose(verbose bool) Opts {
	return func(opt *Options) {
		opt.Verbose = verbose
	}
}
