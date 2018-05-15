package cmd

import (
	true "github.com/jcmdln/cugo/src/true"
	"github.com/spf13/cobra"
)

var (
	trueCmd = &cobra.Command{
		Use:   "true",
		Short: "Return true value",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			true.True()
		},
	}
)

func init() {
	RootCmd.AddCommand(trueCmd)
}
