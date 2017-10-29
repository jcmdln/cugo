package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "Makes directories",
	Long:  `Creates the specified directories if they do not already exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		Mkdir(args...)
	},
}

var (
	mkdirMode    uint32
	mkdirParents bool
	mkdirVerbose bool
)

func init() {
	RootCmd.AddCommand(mkdirCmd)
	mkdirCmd.Flags().Uint32VarP(&mkdirMode, "mode", "m", 0755,
		"Set folder permissions to specified MODE value")
	mkdirCmd.Flags().BoolVarP(&mkdirParents, "parents", "p", false,
		"Create missing parent directories")
	mkdirCmd.Flags().BoolVarP(&mkdirVerbose, "verbose", "v", false,
		"Print a message when actions are taken")
}

func Mkdir(args ...string) {
	for _, dir := range args {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if mkdirParents == false {
				if !strings.Contains(dir, "/") {
					os.Mkdir(dir, os.FileMode(mkdirMode))
					if mkdirVerbose == true {
						fmt.Println("Created directory", "'"+dir+"'",
							"with permissions", os.FileMode(mkdirMode))
					}
				} else {
					fmt.Println("cugo: mkdir:", err)
				}
			} else {
				os.MkdirAll(dir, os.FileMode(mkdirMode))
				if mkdirVerbose == true {
					fmt.Println("Created directory", "'"+dir+"'",
						"with permissions", os.FileMode(mkdirMode))
				}
			}
		}
	}
}
