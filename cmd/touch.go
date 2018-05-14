package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	touchCmd = &cobra.Command{
		Use:   "touch",
		Short: "Change file access and modification times",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Printf("cugo: touch: No operands passed\n" +
					"Usage: touch [-acm] " +
					"[-r REF_FILE|-t TIME|-d DATETIME] TARGETS...\n")
				os.Exit(0)
			} else {
				Touch(args)
			}
		},
	}

	// touchAccess   int
	touchCreate bool
	// touchDate     string
	// touchModified int
	touchVerbose bool
)

func init() {
	RootCmd.AddCommand(touchCmd)
	touchCmd.Flags().SortFlags = false
	// touchCmd.Flags().IntVarP(&touchAccess, "access", "a", 0,
	// 	"Change the access time of a file if -m is also specified")
	touchCmd.Flags().BoolVarP(&touchCreate, "create", "c", false,
		"Do not create the specified file if it doesn't exist")
	// touchCmd.Flags().StringVarP(&touchDate, "date", "d", "",
	// 	"")
	// touchCmd.Flags().IntVarP(&touchModified, "modified", "m", 0,
	// 	"Change the modified time of a file if -a is also specified")
	touchCmd.Flags().BoolVarP(&touchVerbose, "verbose", "v", false,
		"Verbose")
}

func Touch(args []string) {
	Exists := func(t string) bool {
		_, err := os.Stat(t)
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	}

	Verbose := func(t string) {
		if touchVerbose {
			fmt.Printf("cugo: touch: Created %s\n", t)
		}
	}

	for _, target := range args {
		if Exists(target) == false {
			if !touchCreate {
				os.Create(target)
				Verbose(target)
			}
		}
	}
}
