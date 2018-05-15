package cmd

import (
	"fmt"
	"os"

	cugo "github.com/jcmdln/cugo/src/touch"
	"github.com/spf13/cobra"
)

var (
	touchCmd = &cobra.Command{
		Use:   "touch",
		Short: "Change file access and modification times",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Printf("cugo: touch: No operands passed\n" +
					"Usage: touch [-acm] " +
					"[-r REF_FILE|-t TIME|-d DATETIME] TARGETS...\n")
				os.Exit(0)
			} else {
				cugo.Touch(args)
			}
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
