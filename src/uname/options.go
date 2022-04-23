// SPDX-License-Identifier: ISC

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
