package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cugo",
	Short: "Core utilities as a multi-call binary",
	Long:  "cugo is a multi-call binary that provides core utilities for Unix systems.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("cugo: No operands passed\n" +
				"Usage: cugo [COMMAND] [OPTIONS]... ARGUMENTS...")
			os.Exit(0)
		}
	},
}
