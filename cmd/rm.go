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
	Verbose := func(tgt string) {
		if rmVerbose {
			fmt.Println("cugo: rm: removed", tgt)
		}
	}

	Empty := func(name string) bool {
		t, err := os.Open(name)
		defer t.Close()
		_, err = t.Readdirnames(1)
		if err == io.EOF {
			return true
		}
		return false
	}

	for _, target := range args {
		t, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Println("cugo: rm: Can't remove", "'"+target+"':",
				"no such file or directory")
			return
		}

		if !rmRecursive {
			if t.IsDir() {
				fmt.Println("cugo: rm: Can't remove directory '" +
					target + "'")
				return
			}
		}

		if rmForce {
			if rmInteractive || !rmInteractive {
				os.RemoveAll(target)
				Verbose(target)
				return
			}
		}

		if rmRecursive {
			if !t.IsDir() {
				os.Remove(target)
				Verbose(target)
			}

			for !Empty(target) {
				filepath.Walk(target,
					func(t string, info os.FileInfo, err error) error {
						if info.IsDir() && Empty(t) {
							os.Remove(t)
							Verbose(t)
						}

						if !info.IsDir() {
							os.Remove(t)
							Verbose(t)
						}
						return nil
					},
				)
			}

			if Empty(target) {
				os.Remove(target)
			}
		}
	}

	return
}
