package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	basenameCmd = &cobra.Command{
		Use:   "basename",
		Short: "return non-directory portion of a pathname",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
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
