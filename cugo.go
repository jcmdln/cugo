package main

import (
	"github.com/therealfakemoot/cugo/cmd"
	"runtime/debug"
)

func main() {
	debug.SetGCPercent(-1)
	cmd.Execute()
}
