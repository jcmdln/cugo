package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	mkdirCmd = &cobra.Command{
		Use:   "mkdir",
		Short: "Creates the specified directories if they do not already exist",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("cugo: mkdir: No operands passed")
				fmt.Println("Usage: mkdir [-pv] [-m MODE] TARGETS...")
				os.Exit(0)
			} else {
				Mkdir(args)
			}
		},
	}

	mkdirMode    uint32
	mkdirParents bool
	mkdirVerbose bool
)

func init() {
	RootCmd.AddCommand(mkdirCmd)
	mkdirCmd.Flags().Uint32VarP(&mkdirMode, "mode", "m", 0777,
		"Set folder permissions to specified MODE value")
	mkdirCmd.Flags().BoolVarP(&mkdirParents, "parents", "p", false,
		"Create missing parent directories")
	mkdirCmd.Flags().BoolVarP(&mkdirVerbose, "verbose", "v", false,
		"Print a message when actions are taken")
}

func Mkdir(args []string) {
	for _, target := range args {
		_, err := os.Stat(target)
		if os.IsExist(err) {
			fmt.Println("cugo:", "'"+target+"'", "already exists!")
			return
		}

		if !mkdirParents && strings.Contains(target, "/") {
			fmt.Println("cugo: Can't create", "'"+target+"'")
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
