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

package uname

type Options struct {
	All      bool // -a
	Machine  bool // -m
	Nodename bool // -n
	Release  bool // -r
	Sysname  bool // -s
	Version  bool // -v
}

type Opts func(*Options)

func All(all bool) Opts {
	return func(opt *Options) {
		opt.All = all
	}
}

func Machine(machine bool) Opts {
	return func(opt *Options) {
		opt.Machine = machine
	}
}

func Nodename(nodename bool) Opts {
	return func(opt *Options) {
		opt.Nodename = nodename
	}
}

func Release(release bool) Opts {
	return func(opt *Options) {
		opt.Release = release
	}
}

func Sysname(sysname bool) Opts {
	return func(opt *Options) {
		opt.Sysname = sysname
	}
}

func Version(version bool) Opts {
	return func(opt *Options) {
		opt.Version = version
	}
}
