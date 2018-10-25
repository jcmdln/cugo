// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/basename"
	"github.com/spf13/cobra"
)

var (
	basenameCmd = &cobra.Command{
		Use:   "basename",
		Short: "return non-directory portion of a pathname",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			basename.Basename(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(basenameCmd)
}
