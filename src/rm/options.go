// options.go
//
// ISC License
//
// Copyright (c) 2022 Johnathan C. Maudlin <jcmdln@gmail.com>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE
// OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

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
