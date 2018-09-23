package cmd

import (
	"github.com/jcmdln/cugo/src/rmdir"
	"github.com/spf13/cobra"
)

var (
	rmdirCmd = &cobra.Command{
		Use:   "rmdir",
		Short: "Remove directories",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rmdir.Rmdir(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(rmdirCmd)
	rmdirCmd.Flags().SortFlags = false
	rmdirCmd.Flags().BoolVarP(&rmdir.Parents, "parents", "p", false,
		"Remove parent directories")
	rmdirCmd.Flags().BoolVarP(&rmdir.Verbose, "verbose", "v", false,
		"Print a message when actions are taken")
}
