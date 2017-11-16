package cmd

import (
	"github.com/spf13/cobra"
)

var (
	touchCmd = &cobra.Command{
		Use:   "touch",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Touch()
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
	return
}
