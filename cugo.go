package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/jcmdln/cugo/cmd"
)

var (
	Name  = "cugo"
	About = "Core utilities as a multi-call binary written in Go"
	Use   = "[COMMAND] [OPTIONS]... ARGUMENTS..."
)

func init() {
	// Disable GC by default
	debug.SetGCPercent(-1)

	if len(os.Args[1:]) < 1 {
		fmt.Println(Name, "-", About)
		fmt.Println("Usage:", Name, Use)
		os.Exit(0)
	}
}

func main() {
	call := os.Args[1]
	c, ok := cmd.Cmds[call]
	if !ok {
		// Build an array of every utility defined in cmd.Cmds
		utils, u := make([]string, len(cmd.Cmds)), 0
		for i := range cmd.Cmds {
			utils[u] = i
			u++
		}
		fmt.Printf("cugo: '%s' is not a defined utility.\n", call)
		fmt.Printf("Available utilities: %s\n", strings.Join(utils, " "))
		os.Exit(0)
	} else {
		// Trigger GC before running the selected utility
		runtime.GC()

		c.Init(os.Args[2:])
		retVal := c.Main()
		os.Exit(retVal)
	}
}
