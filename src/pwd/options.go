// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package pwd

type Options struct {
	L bool
	P bool
}

type Opts func(*Options)

func L(l bool) Opts {
	return func(opt *Options) {
		opt.L = l
	}
}

func P(p bool) Opts {
	return func(opt *Options) {
		opt.P = p
	}
}
