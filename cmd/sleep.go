package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	sleepCmd = &cobra.Command{
		Use:   "sleep",
		Short: "delay for a specified amount of time",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Sleep(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(sleepCmd)
}

func Sleep(args []string) {
	for _, i := range args {
		t, _ := time.ParseDuration(i)

		time.Sleep(t)
	}
}
