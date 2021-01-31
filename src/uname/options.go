// options.go
//
// Copyright 2021 Johnathan C. Maudlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

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
