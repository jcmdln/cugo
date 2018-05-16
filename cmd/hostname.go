package cmd

import (
	"github.com/jcmdln/cugo/src/hostname"
	"github.com/spf13/cobra"
)

var (
	hostnameCmd = &cobra.Command{
		Use:   "hostname",
		Short: "Return the host name reported by the kernel",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			hostname.Hostname(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(hostnameCmd)
}
