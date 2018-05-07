package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	pwdCmd = &cobra.Command{
		Use:   "pwd",
		Short: "return working directory name",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Pwd()
		},
	}

	pwdL bool
	pwdP bool
)

func init() {
	RootCmd.AddCommand(pwdCmd)
}

func Pwd() {
	p, _ := os.Getwd()
	fmt.Printf("%s\n", p)
}
