// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/hostname"
	"github.com/spf13/cobra"
)

var (
	hostnameCmd = &cobra.Command{
		Use:   "hostname",
		Short: "return the host name reported by the kernel",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			hostname.Hostname(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(hostnameCmd)
	hostnameCmd.Flags().BoolVarP(&hostname.Strip, "strip", "s", false,
		"Trim any domain information from the printed name")
}
