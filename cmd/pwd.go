// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/pwd"
	"github.com/spf13/cobra"
)

var (
	pwdCmd = &cobra.Command{
		Use:   "pwd",
		Short: "return working directory name",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			pwd.Pwd()
		},
	}
)

func init() {
	RootCmd.AddCommand(pwdCmd)
	pwdCmd.Flags().BoolVarP(&pwd.L, "logical", "L", false,
		"Read current dir from env, even symlinks")
	pwdCmd.Flags().BoolVarP(&pwd.P, "physical", "P", false,
		"Absolute path of current dir without symlinks")
}
