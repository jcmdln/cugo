package cmd

import (
	"fmt"
	"io"
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
			if len(args) < 1 {
				fmt.Printf("cugo: rm: No operands passed\n" +
					"Usage: rm [-f|-i] [-r] TARGETS ...\n")
				os.Exit(0)
			} else {
				Rm(args)
			}
		},
	}

	rmDir         bool
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
	rmCmd.Flags().BoolVarP(&rmDir, "dir", "d", false,
		"Remove empty directories")
	rmCmd.Flags().BoolVarP(&rmInteractive, "interactive", "i", false,
		"Prompt before each removal")
	rmCmd.Flags().BoolVarP(&rmRecursive, "recursive", "r", false,
		"Remove directories and their contents recursively")
	rmCmd.Flags().BoolVarP(&rmVerbose, "verbose", "v", false,
		"Print a message when actions are taken")
}

func Rm(args []string) {
	Empty := func(name string) bool {
		t, err := os.Open(name)
		defer t.Close()
		_, err = t.Readdirnames(1)
		if err == io.EOF {
			return true
		}
		return false
	}

	Prompt := func(text string) bool {
		if rmInteractive {
			fmt.Printf(text + " [Yes/No]: ")

			var a string
			_, err := fmt.Scan(&a)
			if err != nil {
				fmt.Println(err)
				return false
			}

			a = strings.ToLower(a)
			if a == "y" || a == "yes" {
				return true
			}
			return false
		} else {
			return true
		}
	}

	Verbose := func(tgt string) {
		if rmVerbose {
			fmt.Println("cugo: rm: removed", tgt)
		}
	}

	for _, target := range args {
		t, err := os.Stat(target)
		if os.IsNotExist(err) {
			fmt.Println("cugo: rm: Can't remove", "'"+target+"':",
				"no such file or directory")
			return
		}

		if rmForce {
			if t.IsDir() {
				fmt.Println("cugo: rm: Can't remove directory '" +
					target + "'")
				return
			} else {
				os.RemoveAll(target)
				Verbose(target)
				return
			}
		}

		if t.IsDir() {
			if rmDir && Empty(target) {
				if Prompt("Remove '" + target + "'?") {
					os.Remove(target)
					Verbose(target)
				}
				return
			}

			if rmRecursive {
				for !Empty(target) {
					filepath.Walk(target,
						func(t string, info os.FileInfo, err error) error {
							if info.IsDir() && Empty(t) {
								if Prompt("Remove '" + t + "'?") {
									os.Remove(t)
									Verbose(t)
								}
							}

							if !info.IsDir() {
								if Prompt("Remove '" + t + "'?") {
									os.Remove(t)
									Verbose(t)
								}
							}

							return nil
						},
					)
				}

				if Empty(target) {
					if Prompt("Remove '" + target + "'?") {
						os.Remove(target)
						Verbose(target)
					}
				}
			} else {
				fmt.Println("cugo: rm: Can't remove directory '" +
					target + "'")
				return
			}
		} else {
			if Prompt("Remove '" + target + "'?") {
				os.Remove(target)
				Verbose(target)
			}
		}
	}

	return
}
