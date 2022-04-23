// SPDX-License-Identifier: ISC

package hostname

type Options struct {
	Strip bool
}

type Opts func(*Options)

func Strip(strip bool) Opts {
	return func(opt *Options) {
		opt.Strip = strip
	}
}
