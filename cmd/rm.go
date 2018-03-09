package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	rmForce       bool
	rmInteractive bool
	rmRecursive   bool
	rmVerbose     bool
)

func init() {
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
	)

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

		if !rmRecursive && t.IsDir() {
			fmt.Println("cugo: Can't remove directory '" + target + "'")
			return
		}

		if rmForce {
			if !rmInteractive || rmInteractive {
				os.RemoveAll(target)
				return
			}
		}

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
