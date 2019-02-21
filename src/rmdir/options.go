// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
