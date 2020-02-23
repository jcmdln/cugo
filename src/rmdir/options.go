// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

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
