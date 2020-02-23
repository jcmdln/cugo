// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package chmod

type Options struct {
	Recursive bool
}

type Opts func(*Options)
