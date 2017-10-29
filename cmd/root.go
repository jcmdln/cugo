package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cugo",
	Short: "Core utilities written in Go",
	Long:  `Cugo is an implementation of standards-compliant core utilities for Unix-like systems in the form of a multi-call binary.`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
