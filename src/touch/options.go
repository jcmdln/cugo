// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package touch

type Options struct {
	Access    bool
	Create    bool
	Date      string
	Modified  bool
	Reference string
}

type Opts func(*Options)

func Access(access bool) Opts {
	return func(opt *Options) {
		opt.Access = access
	}
}

func Create(create bool) Opts {
	return func(opt *Options) {
		opt.Create = create
	}
}

func Date(Date string) Opts {
	return func(opt *Options) {
		opt.Date = Date
	}
}

func Modified(modified bool) Opts {
	return func(opt *Options) {
		opt.Modified = modified
	}
}

func Reference(reference string) Opts {
	return func(opt *Options) {
		opt.Reference = reference
	}
}
