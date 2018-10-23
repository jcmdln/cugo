// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/src/touch"
	"github.com/spf13/cobra"
)

var (
	touchCmd = &cobra.Command{
		Use:   "touch",
		Short: "Change file access and modification times",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			touch.Touch(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(touchCmd)
	touchCmd.Flags().SortFlags = false
	touchCmd.Flags().BoolVarP(&touch.Access, "access", "a", false,
		"Change the access time")
	touchCmd.Flags().BoolVarP(&touch.Create, "create", "c", true,
		"Do not create missing files")
	touchCmd.Flags().StringVarP(&touch.Date, "date", "d", "",
		"Change access and modified time as per ISO8601/RFC3339Nano")
	touchCmd.Flags().BoolVarP(&touch.Modified, "modified", "m", false,
		"Change the modified time")
	touchCmd.Flags().StringVarP(&touch.Reference, "reference", "r", "",
		"Use access and modified time from reference file")
}
