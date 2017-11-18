package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	touchCmd = &cobra.Command{
		Use:   "touch",
		Short: "Change file access and modification times",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("cugo: touch: No operands passed")
				fmt.Println("Usage: touch [-acm] [-r ref_file|-t time|-d date_time] TARGETS...")
				os.Exit(0)
			} else {
				Touch()
			}
		},
	}

	touchAccess   bool
	touchCreate   bool
	touchDate     string
	touchModified int
)

func init() {
	RootCmd.AddCommand(touchCmd)
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
	fmt.Println("cugo: touch: Sorry though 'touch' is currently being implemented.")
	return
}
