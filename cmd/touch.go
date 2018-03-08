package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	touchAccess   bool
	touchCreate   bool
	touchDate     string
	touchModified int
)

func init() {
	var (
		touchCmd = &cobra.Command{
			Use:   "touch",
			Short: "Change file access and modification times",
			Long:  "",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) < 1 {
					fmt.Printf("cugo: touch: No operands passed\n" +
						"Usage: touch [-acm] " +
						"[-r ref_file|-t time|-d date_time] TARGETS...")
					os.Exit(0)
				} else {
					Touch()
				}
			},
		}
	)

	RootCmd.AddCommand(touchCmd)
	touchCmd.Flags().SortFlags = false
	touchCmd.Flags().BoolVarP(&touchAccess, "access", "a", false,
		"Change the access time of a file if -m is also specified")
	touchCmd.Flags().BoolVarP(&touchCreate, "create", "c", false,
		"Do not create the specified file if it doesn't exist")
	touchCmd.Flags().StringVarP(&touchDate, "date", "d", "",
		"")
	touchCmd.Flags().IntVarP(&touchModified, "modified", "m", 0,
		"Changed the modification time of a file if -a is also specified")
}

func Touch() {
	fmt.Println("cugo: touch: Not yet implemented.")
	return
}
