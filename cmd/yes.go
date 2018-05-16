package cmd

import (
	"github.com/jcmdln/cugo/src/yes"
	"github.com/spf13/cobra"
)

var (
	yesCmd = &cobra.Command{
		Use:   "yes",
		Short: "Repeatedly output specified string(s), or 'y'",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			yes.Yes(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(yesCmd)
}
