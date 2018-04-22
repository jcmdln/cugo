package cmd

import (
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	sleepCmd = &cobra.Command{
		Use:   "sleep",
		Short: "delay for a specified amount of time",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Sleep(args)
		},
	}
)

func init() {
	RootCmd.AddCommand(sleepCmd)
}

func Sleep(args []string) {
	Time := func(in string, txt string) {
		if strings.Contains(in, txt) {
			s := strings.Split(in, txt)
			t, _ := strconv.Atoi(s[0])
			if txt == "" || txt == "s" {
				time.Sleep(time.Duration(t) * time.Second)
			}
			if txt == "h" {
				time.Sleep(time.Duration(t) * time.Minute)
			}
			if txt == "m" {
				time.Sleep(time.Duration(t) * time.Hour)
			}
			if txt == "d" {
				time.Sleep(time.Duration(t) * time.Hour * 24)
			}
		}
	}

	for _, i := range args {
		Time(i, "")
		Time(i, "s")
		Time(i, "m")
		Time(i, "h")
		Time(i, "d")
	}

	return
}
