// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

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
