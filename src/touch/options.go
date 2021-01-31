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
