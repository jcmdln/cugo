// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cat

type Options struct {
	Unbuffered bool
}

type Opts func(*Options)

func Unbuffered(unbuffered bool) Opts {
	return func(opts *Options) {
		opts.Unbuffered = unbuffered
	}
}
