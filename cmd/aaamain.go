// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package cmd provides the command line interface for `cugo`
package cmd

import (
	"github.com/jcmdln/flagger/commands"
)

var Command *commands.Commands

func init() {
	Command = commands.New()
}
