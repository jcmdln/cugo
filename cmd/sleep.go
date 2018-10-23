// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"strings"

	"github.com/jcmdln/cugo/src/sleep"
	"github.com/spf13/cobra"
)

var (
	sleepCmd = &cobra.Command{
		Use:   "sleep",
		Short: "Delay for a specified amount of time",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			sleep.Sleep(strings.Join(args, " "))
		},
	}
)

func init() {
	RootCmd.AddCommand(sleepCmd)
}
