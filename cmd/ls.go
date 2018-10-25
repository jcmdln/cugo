// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/ls"
	"github.com/spf13/cobra"
)

var (
	lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "list files and directories",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			ls.Ls(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(lsCmd)
	lsCmd.Flags().SortFlags = false
	lsCmd.Flags().BoolVarP(&ls.All, "all", "a", false,
		"Include directory entries that begin with '.'")
	lsCmd.Flags().BoolVarP(&ls.Recursive, "recursive", "r", false,
		"Recursively traverse folders")
}
