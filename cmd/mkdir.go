package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	mkdirMode    uint32
	mkdirParents bool
	mkdirVerbose bool
)

func init() {
	var (
		mkdirCmd = &cobra.Command{
			Use:   "mkdir",
			Short: "Create directories",
			Long:  "",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) < 1 {
					fmt.Printf("cugo: mkdir: No operands passed\n" +
						"Usage: mkdir [-pv] [-m MODE] DIRECTORIES ...")
					os.Exit(0)
				} else {
					Mkdir(args)
				}
			},
		}
	)

	RootCmd.AddCommand(mkdirCmd)
	mkdirCmd.Flags().SortFlags = false
	mkdirCmd.Flags().Uint32VarP(&mkdirMode, "mode", "m", 0777,
		"Set directory permissions to MODE value")
	mkdirCmd.Flags().BoolVarP(&mkdirParents, "parents", "p", false,
		"Create missing parent directories")
	mkdirCmd.Flags().BoolVarP(&mkdirVerbose, "verbose", "v", false,
		"Verbose")
}

func Mkdir(args []string) {
	for _, target := range args {
		_, err := os.Stat(target)
		if os.IsExist(err) {
			fmt.Println("cugo: mkdir:", "'"+target+"'",
				"already exists!")
			return
		}

		if !mkdirParents && strings.Contains(target, "/") {
			fmt.Println("cugo: mkdir: Can't create", "'"+target+"'")
		} else if mkdirParents == true {
			os.MkdirAll(target, os.FileMode(mkdirMode))
		}

		if mkdirVerbose == true {
			fmt.Println("Created", "'"+target+"'", "with permissions",
				os.FileMode(mkdirMode))
		}
	}

	return
}
