package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	pwdCmd = &cobra.Command{
		Use:   "pwd",
		Short: "Return working directory name",
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
	pwdCmd.Flags().BoolVarP(&pwdL, "logical", "L", false,
		"Read current dir from env, even symlinks")
	pwdCmd.Flags().BoolVarP(&pwdP, "physical", "P", false,
		"Absolute path of current dir without symlinks")
}

func Pwd() {
	var dir string

	if pwdP {
		d, _ := os.Getwd()
		dir, _ = filepath.EvalSymlinks(d)
	} else {
		dir = os.Getenv("PWD")
	}

	fmt.Printf("%s\n", dir)
}
