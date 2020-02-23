// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package ls

type Options struct {
	All       bool
	Recursive bool
}

type Opts func(*Options)
