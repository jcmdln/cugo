package cmd

import (
	"github.com/jcmdln/cugo/src/ls"
	"github.com/spf13/cobra"
)

var (
	lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "List files and directories",
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
		"")
	lsCmd.Flags().BoolVarP(&ls.Interactive, "interactive", "i", false,
		"")
	lsCmd.Flags().BoolVarP(&ls.Recursive, "recursive", "r", false,
		"")
	lsCmd.Flags().BoolVarP(&ls.Verbose, "verbose", "v", false,
		"")
}
