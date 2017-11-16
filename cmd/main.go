package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cugo",
	Short: "Core utilities as a multi-call binary written in Go",
	Long:  "Cugo is a set of core utilities for Unix-like systems in the form of a multi-call binary with the aim of being broadly standards compliant.",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}
