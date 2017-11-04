package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/jcmdln/cugo/cmd"
)

var (
	Name  = "cugo"
	Use   = "[COMMAND] [OPTIONS]... ARGUMENTS..."
	Utils = "mkdir"
)

func init() {
	debug.SetGCPercent(-1)
	if len(os.Args[1:]) < 1 {
		fmt.Println(Name+":", "not enough operands")
		fmt.Println(Name+":", Use)
		os.Exit(0)
	}
}

func main() {
	call := os.Args[1]
	c, ok := cmd.Cmds[call]
	if !ok {
		fmt.Printf("cugo: '%s' is not a defined utility.\n", call)
		fmt.Printf("Available utilities: %s\n", Utils)
		os.Exit(0)
	} else {
		runtime.GC()
		c.Init(os.Args[2:])
		retVal := c.Main()
		os.Exit(retVal)
	}
}
