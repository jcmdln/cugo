package cmd

import (
	cugo "github.com/jcmdln/cugo/src/count"
	"github.com/spf13/cobra"
)

var (
	countCmd = &cobra.Command{
		Use:   "count",
		Short: "Count the number of elements of an array",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cugo.Count(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(countCmd)
}
