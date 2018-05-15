package cmd

import (
	cugo "github.com/jcmdln/cugo/src/basename"
	"github.com/spf13/cobra"
)

var (
	basenameCmd = &cobra.Command{
		Use:   "basename",
		Short: "Return non-directory portion of a pathname",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cugo.Basename(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(basenameCmd)
}
