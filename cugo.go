package main

import (
	"runtime/debug"

	"github.com/jcmdln/cugo/cmd"
)

func init() {
	debug.SetGCPercent(-1)
}

func main() {
	cmd.RootCmd.Execute()
}
