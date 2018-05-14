package cmd

import (
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	sleepCmd = &cobra.Command{
		Use:   "sleep",
		Short: "Delay for a specified amount of time",
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
		s := strings.Split(i, "")
		if len(s) < 2 {
			i += "s"
		}
		t, _ := time.ParseDuration(i)
		time.Sleep(t)
	}
}
