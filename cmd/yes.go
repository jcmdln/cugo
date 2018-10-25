// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/yes"
	"github.com/spf13/cobra"
)

var (
	yesCmd = &cobra.Command{
		Use:   "yes",
		Short: "repeatedly output specified string(s), or 'y'",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			yes.Yes(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(yesCmd)
}
