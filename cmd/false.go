package cmd

import (
	"github.com/spf13/cobra"
)

var (
	falseCmd = &cobra.Command{
		Use:   "false",
		Short: "Return false value",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			False()
		},
	}
)

func init() {
	RootCmd.AddCommand(falseCmd)
}

func False() bool {
	return false
}
