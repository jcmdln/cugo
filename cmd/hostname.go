package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	hostnameCmd = &cobra.Command{
		Use:   "hostname",
		Short: "return the host name reported by the kernel",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Hostname(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(hostnameCmd)
}

func Hostname(args []string) {
	h, _ := os.Hostname()
	fmt.Printf("%s\n", h)
}
