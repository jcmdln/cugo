package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	rmCmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove directory entries",
		Long:  "Remove the directory entry specified by each file argument",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Printf("cugo: rm: No operands passed\n" +
					"Usage: rm [-f|-i] [-r] TARGETS ...\n")
				os.Exit(0)
			} else {
				Rm(args)
			}
		},
	}

	rmForce       bool
	rmInteractive bool
	rmRecursive   bool
	rmVerbose     bool
)

func init() {
	RootCmd.AddCommand(rmCmd)
	rmCmd.Flags().SortFlags = false
	rmCmd.Flags().BoolVarP(&rmForce, "force", "f", false,
		"Skip prompts and ignore warnings")
	rmCmd.Flags().BoolVarP(&rmInteractive, "interactive", "i", false,
		"Prompt before performing each action")
	rmCmd.Flags().BoolVarP(&rmRecursive, "recursive", "r", false,
		"Remove directories and their contents recursively")
	rmCmd.Flags().BoolVarP(&rmVerbose, "verbose", "v", false,
		"Print a message when actions are taken")
}

func Rm(args []string) {
	for _, target := range args {
		t, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Println("cugo: rm: Cannot remove", "'"+target+"':",
				"no such file or directory")
			return
		}

		if !rmRecursive && t.IsDir() {
			fmt.Println("cugo: rm: Can't remove directory '" +
				target + "'")
			return
		}

		if rmForce {
			if !rmInteractive || rmInteractive {
				//os.RemoveAll(target)
				return
			}
		}

		if rmRecursive && !rmForce {
			filepath.Walk(target,
				func(path string, info os.FileInfo, err error) error {
					if info.IsDir() {

						fmt.Println("dir :", path)
					} else {
						fmt.Println("file:", path)
					}
					return nil
				},
			)
		}

		if rmVerbose {
			fmt.Println("cugo: Removed '" + target + "'")
		}
	}

	return
}
