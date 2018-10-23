// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/false"
	"github.com/spf13/cobra"
)

var (
	falseCmd = &cobra.Command{
		Use:   "false",
		Short: "Return false value",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			false.False()
		},
	}
)

func init() {
	RootCmd.AddCommand(falseCmd)
}
