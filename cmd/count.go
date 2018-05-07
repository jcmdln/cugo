package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	countCmd = &cobra.Command{
		Use:   "count",
		Short: "count the number of elements of an array",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Count(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(countCmd)
}

func Count(args []string) {
	var total int
	for i, _ := range args {
		total = i + 1
	}
	fmt.Printf("%v\n", total)
}
