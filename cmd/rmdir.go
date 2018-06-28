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
	rmCmd.Flags().SortFlags = false
	rmCmd.Flags().BoolVarP(&rmdir.Pathname, "pathname", "p", false,
		"Skip prompts and ignore warnings")
}
