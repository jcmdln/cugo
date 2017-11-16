package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	rmCmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove directory entries",
		Long:  "Remove the directory entry specified by each file argument",
		Run: func(cmd *cobra.Command, args []string) {
			if len(os.Args) <= 2 {
				fmt.Println("cugo: rm: No operands passed")
				fmt.Println("Usage: rm [-f|-i] [-r] TARGETS...")
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
		_, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Println("cugo:", "Cannot remove '"+target+"':",
				"no such file or directory")
			return
		}

		// '-i' is ignored if '-f' is passed
		if rmInteractive == true && !rmForce {
			Read([]string{"cugo: Remove " + target + "? [Y/N]: "})
		}

		// '-r' walks the path if '-f' is not passed to the inner-most target.
		if rmRecursive == true && !rmForce {
			filepath.Walk(target,
				func(t string, info os.FileInfo, err error) error {
					if info.IsDir() {
						Read([]string{"cugo: Descend into '" + info.Name() + "'?: "})
						return nil
					} else if !info.IsDir() {
						Read([]string{"cugo: Delete '" + info.Name() + "'?: "})
						fmt.Println("pretend to delete", t)
						return nil
					}
					fmt.Println("pretend to delete", t)
					return nil
				})
		}

		// '-f' cannot recursively descend unless '-r' is passed
		if rmForce == true && !rmRecursive && !strings.Contains(target, "/") {
			fmt.Println("pretend to delete", target)
		} else if strings.Contains(target, "/") {
			fmt.Println("cugo:", "Cannot remove the target", "'"+target+"'")
		}

		if rmVerbose == true {
			fmt.Println("cugo:", "Removed '"+target+"'")
		}
	}

	return
}
