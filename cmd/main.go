package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cugo",
	Short: "Core utilities as a multi-call binary written in Go",
	Long:  "Cugo is a set of core utilities for Unix-like systems in the form of a multi-call binary with the aim of being broadly standards compliant.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("cugo: No operands passed")
			fmt.Println("Usage: cugo [COMMAND] [OPTIONS]... ARGUMENTS...")
			os.Exit(0)
		} else {
			Cugo()
		}
	},
}

func init() {
	//
}

func Cugo() {
	//
}
