package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

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

	mkdirMode    uint32
	mkdirParents bool
	mkdirVerbose bool
)

func init() {
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
	Exists := func(t string) bool {
		_, err := os.Stat(t)
		if os.IsNotExist(err) {
			return false
		}
		return true
	}

	Verbose := func(t string) {
		fmt.Printf("cugo: mkdir: Created %s\n", t)
	}

	for _, target := range args {
		if Exists(target) {
			fmt.Println("cugo: mkdir: '" + target + "' already exists!")
			return
		}

		if mkdirParents {
			c := "."
			t := strings.Split(filepath.Clean(target), "/")
			for i := range t {
				c += "/" + t[i]
				if !Exists(c) {
					os.Mkdir(c, os.FileMode(mkdirMode))
					Verbose(c)
				}
			}
		} else {
			if !Exists(target) {
				os.Mkdir(target, os.FileMode(mkdirMode))
				Verbose(target)
			}
		}
	}

	return
}
