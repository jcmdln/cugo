// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/count"
	"github.com/spf13/cobra"
)

var (
	countCmd = &cobra.Command{
		Use:   "count",
		Short: "Count the number of elements of an array",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			count.Count(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(countCmd)
}
