package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	basenameCmd = &cobra.Command{
		Use:   "basename",
		Short: "Return non-directory portion of a pathname",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Printf("cugo: basename: No arguments passed\n" +
					"Usage: basename ARGUMENTS...\n")
				return
			}

			Basename(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(basenameCmd)
}

func Basename(args []string) {
	b := filepath.Base(args[0])
	fmt.Printf("%s\n", b)
}
