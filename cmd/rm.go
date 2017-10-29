package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes specified file(s)",
	Long:  `Removes specified file(s) and optionally removes directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		Rm(args...)
	},
}

var (
	RmForce       bool
	RmInteractive bool
	RmRecursive   bool
	RmVerbose     bool
)

func init() {
	RootCmd.AddCommand(rmCmd)
	rmCmd.Flags().BoolVarP(&RmForce, "force", "f", false,
		"Skip prompts and ignore warnings")
	rmCmd.Flags().BoolVarP(&RmInteractive, "interactive", "i", false,
		"Prompt before performing each action")
	rmCmd.Flags().BoolVarP(&RmRecursive, "recursive", "r", false,
		"Remove directories and their contents recursively")
	rmCmd.Flags().BoolVarP(&RmVerbose, "verbose", "v", false,
		"Print a message when actions are taken")
}

func Rm(args ...string) {
	for _, target := range args {
		if _, err := os.Stat(target); err == nil {
			if RmInteractive == true {
			}
			//os.RemoveAll(target)
			if RmVerbose == true {
				fmt.Println("Deleting", target)
			}
		}
	}
}
