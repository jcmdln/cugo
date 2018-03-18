package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var (
	lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "List files and directories",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Printf("cugo: ls: No operands passed\n" +
					"Usage: ls [-i] [-r] TARGETS ...\n")
				os.Exit(0)
			} else {
				Ls(args)
			}
		},
	}

	lsInteractive bool
	lsRecursive   bool
	lsVerbose     bool
)

func init() {
	RootCmd.AddCommand(lsCmd)
	lsCmd.Flags().SortFlags = false
	lsCmd.Flags().BoolVarP(&lsInteractive, "interactive", "i", false,
		"")
	lsCmd.Flags().BoolVarP(&lsRecursive, "recursive", "r", false,
		"")
	lsCmd.Flags().BoolVarP(&lsVerbose, "verbose", "v", false,
		"")
}

func Ls(args []string) {
	for _, target := range args {
		_, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Println("cugo: ls: '"+target+"':",
				"no such file or directory")
			return
		}

		items, err := ioutil.ReadDir(target)
		if err != nil {
			fmt.Println("cugo: rm:", err)
		}

		for _, item := range items {
			fmt.Printf(item.Name() + " ")
		}

		fmt.Printf("\n")
	}

	return
}
