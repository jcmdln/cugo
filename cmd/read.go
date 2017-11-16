package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	readCmd = &cobra.Command{
		Use:   "read",
		Short: "read from standard input into shell variables",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			//
		},
	}
)

func init() {
	RootCmd.AddCommand(readCmd)
}

func Read(msg string) {
	// Read still needs a lengthy cleanup
	// if len(os.Args) <= 2 {
	// 	fmt.Println("cugo: mkdir: No operands passed")
	// 	fmt.Println("Usage: mkdir [-pv] [-m MODE] TARGETS...")
	//  os.Exit(0)
	// }

	fmt.Printf(msg)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		switch strings.ToLower(input.Text()) {
		case "yes", "y":
			return
		case "no", "n":
			os.Exit(0)
		default:
			fmt.Println("cugo:", "'"+input.Text()+"'", "is invalid.")
			fmt.Printf("Please answer 'y|yes' or 'n|no': ")
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(0)
	}

	return
}
