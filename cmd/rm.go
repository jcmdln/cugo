package cmd

import (
	"fmt"
	"io"
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

func isEmpty(name string) (bool, error) {
	target, err := os.Open(name)
	defer target.Close()

	_, err = target.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func Rm(args []string) {
	for _, target := range args {
		t, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Println("cugo:", "Cannot remove '"+target+"':",
				"no such file or directory")
			return
		}

		// Exit if '-r' isn't passed and the target is a directory
		if !rmRecursive && t.IsDir() {
			fmt.Println("cugo: Can't remove directory '" + target + "'")
			return
		}

		// '-i' is ignored if '-f' is passed
		if rmForce {
			if !rmInteractive || rmInteractive {
				os.RemoveAll(target)
				return
			}
		}

		// '-r' walks the path if '-f' is not passed.
		if rmRecursive && !rmForce {
			filepath.Walk(target,
				func(t string, info os.FileInfo, err error) error {
					empty, _ := isEmpty(t)

					if info.IsDir() {
						if !empty {
							Read([]string{"cugo: Descend into '" +
								info.Name() + "'?: "})
						} else if empty {
							Read([]string{"cugo: Remove directory '" +
								info.Name() + "'?: "})
							os.Remove(t)
						}
					} else if !info.IsDir() {
						Read([]string{"cugo: Remove file '" +
							info.Name() + "'?: "})
						os.Remove(t)
					}

					return nil
				})
		}

		if rmVerbose {
			fmt.Println("cugo: Removed '" + target + "'")
		}
	}

	return
}
