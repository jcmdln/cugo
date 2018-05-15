package cmd

import (
	cugo "github.com/jcmdln/cugo/src/mkdir"
	"github.com/spf13/cobra"
)

var (
	mkdirCmd = &cobra.Command{
		Use:   "mkdir",
		Short: "Create directories",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cugo.Mkdir(args)
		},
	}
)

type mkdir struct {
}

func init() {
	mkdir := cugo.MKDIR{}

	RootCmd.AddCommand(mkdirCmd)
	mkdirCmd.Flags().SortFlags = false
	mkdirCmd.Flags().Uint32VarP(&mkdir.Mode, "mode", "m", 0777,
		"Set directory permissions to MODE value")
	mkdirCmd.Flags().BoolVarP(&mkdir.Parents, "parents", "p", false,
		"Create missing parent directories")
	mkdirCmd.Flags().BoolVarP(&mkdir.Verbose, "verbose", "v", false,
		"Verbose")
}
