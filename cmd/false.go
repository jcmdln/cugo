package cmd

import (
	false "github.com/jcmdln/cugo/src/false"
	"github.com/spf13/cobra"
)

var (
	falseCmd = &cobra.Command{
		Use:   "false",
		Short: "Return false value",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			false.False()
		},
	}
)

func init() {
	RootCmd.AddCommand(falseCmd)
}
