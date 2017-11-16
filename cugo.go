package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/jcmdln/cugo-cobra/cmd"
)

func init() {
	debug.SetGCPercent(-1)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
