package cmd

import (
	cugo "github.com/jcmdln/cugo/src/pwd"
	"github.com/spf13/cobra"
)

var (
	pwdCmd = &cobra.Command{
		Use:   "pwd",
		Short: "Return working directory name",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cugo.Pwd()
		},
	}
)

func init() {
	pwd := &cugo.PWD{}

	RootCmd.AddCommand(pwdCmd)
	pwdCmd.Flags().BoolVarP(&pwd.L, "logical", "L", false,
		"Read current dir from env, even symlinks")
	pwdCmd.Flags().BoolVarP(&pwd.P, "physical", "P", false,
		"Absolute path of current dir without symlinks")
}
