package cmd

import (
	"github.com/spf13/cobra"
)

var (
	trueCmd = &cobra.Command{
		Use:   "true",
		Short: "Return true value",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			True()
		},
	}
)

func init() {
	RootCmd.AddCommand(trueCmd)
}

func True() bool {
	return true
}
