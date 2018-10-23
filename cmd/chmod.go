// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/chmod"
	"github.com/spf13/cobra"
)

var (
	chmodCmd = &cobra.Command{
		Use:   "chmod",
		Short: "Change file mode bits",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			chmod.Chmod(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(chmodCmd)
	chmodCmd.Flags().SortFlags = false
	chmodCmd.Flags().BoolVarP(&chmod.Changes, "changes", "c", false,
		"Report only when changes are made")
	chmodCmd.Flags().BoolVarP(&chmod.Quiet, "quiet", "q", false,
		"Suppress most error messages")
	chmodCmd.Flags().BoolVarP(&chmod.Verbose, "verbose", "v", false,
		"Report each processed file")
	chmodCmd.Flags().BoolVarP(&chmod.Recursive, "recursive", "r", false,
		"Change files and directories recursively")
}
