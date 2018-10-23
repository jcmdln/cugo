// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/true"
	"github.com/spf13/cobra"
)

var (
	trueCmd = &cobra.Command{
		Use:   "true",
		Short: "Return true value",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			true.True()
		},
	}
)

func init() {
	RootCmd.AddCommand(trueCmd)
}
