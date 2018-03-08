package cmd

import (
	// "fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	var (
		yesCmd = &cobra.Command{
			Use:   "yes",
			Short: "Repeatedly output specified string(s), or 'y'",
			Long:  "",
			Run: func(cmd *cobra.Command, args []string) {
				Yes(args)
			},
		}
	)

	RootCmd.AddCommand(yesCmd)
}

func Yes(args []string) {
	if len(args) == 0 {
		for true {
			io.WriteString(os.Stdout, "y\n")
		}
	} else {
		for true {
			out := strings.Join(args, " ") + "\n"
			io.WriteString(os.Stdout, out)
		}
	}

	return
}
