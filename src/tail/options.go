// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build todo

package tail

type Options struct {
	Count  int
	Follow bool
	Number int
}

type Opts func(*Options)
