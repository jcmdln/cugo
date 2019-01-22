// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package cmd provides the command line interface for `cugo`
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// RootCmd is the definition of our command, in this case 'cugo'.
	RootCmd = &cobra.Command{
		Use:   "cugo",
		Short: "",
		Long:  "Core Utilities in multi-call Go binary",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)
