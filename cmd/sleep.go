package cmd

import (
	cugo "github.com/jcmdln/cugo/src/sleep"
	"github.com/spf13/cobra"
)

var (
	sleepCmd = &cobra.Command{
		Use:   "sleep",
		Short: "Delay for a specified amount of time",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cugo.Sleep(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(sleepCmd)
}
