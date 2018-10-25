// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/whoami"
	"github.com/spf13/cobra"
)

var (
	whoamiCmd = &cobra.Command{
		Use:   "whoami",
		Short: "return current user",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			whoami.Whoami()
		},
	}
)

func init() {
	RootCmd.AddCommand(whoamiCmd)
}
