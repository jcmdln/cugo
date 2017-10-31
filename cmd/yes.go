package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var yesCmd = &cobra.Command{
	Use:   "yes",
	Short: "Repeatedly echoes a string",
	Long:  `Prints a string to stdout until terminated`,
	Run: func(cmd *cobra.Command, args []string) {
		Yes(args...)
	},
}

func init() {
	RootCmd.AddCommand(yesCmd)
}

// Yes will output the provided string, or "yes", indefinitely.
func Yes(args ...string) {
	var out string

	if len(args) < 1 {
		out = "y"
	} else {
		out = strings.Join(args, " ")
	}

	out = out + "\n"

	for {
		os.Stdout.WriteString(out)
	}
}
