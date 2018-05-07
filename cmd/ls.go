package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "List files and directories",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Ls(args)
		},
	}

	lsAll         bool
	lsInteractive bool
	lsRecursive   bool
	lsVerbose     bool
)

func init() {
	RootCmd.AddCommand(lsCmd)
	lsCmd.Flags().SortFlags = false
	lsCmd.Flags().BoolVarP(&lsAll, "all", "a", false,
		"")
	lsCmd.Flags().BoolVarP(&lsInteractive, "interactive", "i", false,
		"")
	lsCmd.Flags().BoolVarP(&lsRecursive, "recursive", "r", false,
		"")
	lsCmd.Flags().BoolVarP(&lsVerbose, "verbose", "v", false,
		"")
}

func Ls(args []string) {
	List := func(t string) {
		items, err := ioutil.ReadDir(t)
		if err != nil {
			fmt.Println("cugo: rm:", err)
		}
		for _, item := range items {
			if !lsAll && strings.HasPrefix(item.Name(), ".") {
				//
			} else {
				fmt.Printf(item.Name() + " ")
			}
		}
		fmt.Printf("\n")
	}

	if len(args) == 0 {
		List(".")
		return
	}

	for _, target := range args {
		_, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Printf("cugo: ls: '%s': No such file or directory\n", target)
			return
		}
		List(target)
	}

	return
}
