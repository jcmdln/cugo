package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/jcmdln/cugo/cmd"
)

func init() {
	// Disable GC by default
	debug.SetGCPercent(-1)

	// If no flags or args passed, print basic info
	if len(os.Args) <= 1 {
		fmt.Println("cugo", "v0.0.0-alpha", "-",
			"Core utilities as a multi-call binary written in Go")
		fmt.Println("Available commands: mkdir")
		os.Exit(0)
	}
}

func main() {
	call := os.Args[1]
	c, ok := cmd.Cmds[call]
	if !ok {
		fmt.Printf("Unable to call %s\n", call)
		os.Exit(0)
	} else {
		runtime.GC()
		c.Init()
		retVal := c.Main()
		os.Exit(retVal)
	}
}
