package cmd

import (
	cugo "github.com/jcmdln/cugo/src/whoami"
	"github.com/spf13/cobra"
)

var (
	whoamiCmd = &cobra.Command{
		Use:   "whoami",
		Short: "Return current user",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cugo.Whoami()
		},
	}
)

func init() {
	RootCmd.AddCommand(whoamiCmd)
}
