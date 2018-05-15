package cmd

import (
	"fmt"

	cugo "github.com/jcmdln/cugo/src/rm"
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
				return
			} else {
				cugo.Rm(args)
			}
		},
	}
)

func init() {
	rm := &cugo.RM{}

	RootCmd.AddCommand(rmCmd)
	rmCmd.Flags().SortFlags = false
	rmCmd.Flags().BoolVarP(&rm.Force, "force", "f", false,
		"Skip prompts and ignore warnings")
	rmCmd.Flags().BoolVarP(&rm.Dir, "dir", "d", false,
		"Remove empty directories")
	rmCmd.Flags().BoolVarP(&rm.Interactive, "interactive", "i", false,
		"Prompt before each removal")
	rmCmd.Flags().BoolVarP(&rm.Recursive, "recursive", "r", false,
		"Remove directories and their contents recursively")
	rmCmd.Flags().BoolVarP(&rm.Verbose, "verbose", "v", false,
		"Print a message when actions are taken")
}
