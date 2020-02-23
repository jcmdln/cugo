// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

// +build testing

package tail

type Options struct {
	Count  int
	Follow bool
	Number int
}

type Opts func(*Options)
