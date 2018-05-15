package cmd

import (
	cugo "github.com/jcmdln/cugo/src/touch"
	"github.com/spf13/cobra"
)

var (
	touchCmd = &cobra.Command{
		Use:   "touch",
		Short: "Change file access and modification times",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cugo.Touch(args)
		},
	}
)

func init() {
	touch := &cugo.TOUCH{}

	RootCmd.AddCommand(touchCmd)
	touchCmd.Flags().SortFlags = false
	// touchCmd.Flags().IntVarP(&touch.Access, "access", "a", 0,
	// 	"Change the access time of a file if -m is also specified")
	touchCmd.Flags().BoolVarP(&touch.Create, "create", "c", false,
		"Do not create the specified file if it doesn't exist")
	// touchCmd.Flags().StringVarP(&touch.Date, "date", "d", "",
	// 	"")
	// touchCmd.Flags().IntVarP(&touch.Modified, "modified", "m", 0,
	// 	"Change the modified time of a file if -a is also specified")
	touchCmd.Flags().BoolVarP(&touch.Verbose, "verbose", "v", false,
		"Verbose")
}
