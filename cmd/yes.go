package cmd

import (
	// "fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	yesCmd = &cobra.Command{
		Use:   "yes",
		Short: "Echoes 'y' or the provided string until killed.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Yes(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(yesCmd)
}

func Yes(args []string) {
	if len(args) == 0 {
		for true {
			io.WriteString(os.Stdout, "y")
		}
	} else {
		for true {
			io.WriteString(os.Stdout, strings.Join(args, " "))
		}
	}
}
