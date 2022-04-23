// SPDX-License-Identifier: ISC

package head

type Options struct {
	Number int
}

type Opts func(*Options)

func Number(num int) Opts {
	return func(opt *Options) {
		opt.Number = num
	}
}
