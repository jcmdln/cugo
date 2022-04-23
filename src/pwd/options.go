// SPDX-License-Identifier: ISC

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
