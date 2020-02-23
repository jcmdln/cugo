// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

// Package cmd provides the command line interface for `cugo`, courtesy
// of https://github.com/hlfstr/flagger.
package cmd

import (
	"github.com/jcmdln/flagger/commands"
)

var Command *commands.Commands

func init() {
	Command = commands.New()
}
