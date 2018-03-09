package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
						"Usage: mkdir [-pv] [-m MODE] DIRECTORIES ...\n")
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
			fmt.Println("cugo: mkdir: '" + target + "' already exists!")
			return
		}

		if mkdirParents == true {
			os.MkdirAll(target, os.FileMode(mkdirMode))
		} else {
			dir, err := filepath.Abs(target)
			if err == nil {
				fmt.Println("cugo: mkdir: Cannot create '" + dir +
					"', missing parent directory.")
			} else {
				os.Mkdir(dir, os.FileMode(mkdirMode))
			}
		}

		if mkdirVerbose == true {
			fmt.Println("Created '"+target+"' with permissions",
				os.FileMode(mkdirMode))
		}
	}

	return
}
