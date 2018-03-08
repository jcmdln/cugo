package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	readArray    string
	readDelim    string
	readReadline bool
	readText     string
	readnchars   int
	readNchars   int
)

func init() {
	var (
		readCmd = &cobra.Command{
			Use:   "read",
			Short: "read from standard input into shell variables",
			Long:  "",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) < 1 {
					fmt.Printf("cugo: read: No operands passed\n" +
						"Usage: read [-epr] TARGETS...")
					os.Exit(0)
				} else {
					Read(args)
				}
			},
		}
	)

	RootCmd.AddCommand(readCmd)
	readCmd.Flags().SortFlags = false
	readCmd.Flags().StringVar(&readArray, "a", "",
		"Assign read words to sequential indices of ARRAY")
	readCmd.Flags().StringVar(&readDelim, "d", "",
		"Continue until the first character of DELIM is read rather"+
			"than newline")
	readCmd.Flags().BoolVar(&readReadline, "e", false,
		"Use Readline to obtain the line in an interactive shell")
	readCmd.Flags().StringVar(&readText, "i", "",
		"Use TEXT as the initial text for Readline")
	readCmd.Flags().IntVar(&readnchars, "n", 0,
		"Return after reading NCHARS rather than newline, but honor a"+
			"delimiter if fewer than NCHARS are read")
	readCmd.Flags().IntVar(&readNchars, "N", 0,
		"Return after NCHARS unless EOF or timeout is first")

}

func Read(args []string) {

	return
}
