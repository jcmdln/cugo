package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	rmCmd = &cobra.Command{
		Use:   "rm",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Rm(args)
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

		if rmInteractive == true && rmForce == false {
			Read("cugo: Remove " + target + "? [Y/N]: ")
		}

		if rmForce == true {
			os.RemoveAll(target)
		} else if rmRecursive == true && rmForce == false {
			filepath.Walk(target,
				func(t string, info os.FileInfo, err error) error {
					if info.IsDir() {
						Read("cugo: Descend into '" + info.Name() + "'?: ")
						return nil
					}
					fmt.Println("pretend to delete", t)
					return nil
				})
		} else if !rmRecursive && !strings.Contains(target, "/") {
			fmt.Println("cugo:", "Cannot remove the directory",
				"'"+target+"'")
		} else {
			os.Remove(target)
		}

		if rmVerbose == true {
			fmt.Println("cugo:", "Removed '"+target+"'")
		}
	}

	return
}
