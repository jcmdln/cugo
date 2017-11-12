package cmd

import (
	"fmt"
	"os"

	"github.com/markedhero/flagit"
)

type Commander interface {
	Init(flag []string)
	Main() int
}

type Command struct {
	Name  string // Name of the utility
	About string // Short description
	Use   string // Usage legend
	Man   string // Manual page

	Flags *flagit.Flag
	Data  []string
}

func (c *Command) ArgParse(opt int) {
	var err error
	c.Data, err = c.Flags.Parse(os.Args[2:])

	if len(os.Args[opt:]) < 1 {
		fmt.Println(c.Name+":", "Not enough operands")
		fmt.Println("Usage:", c.Name, c.Use)
		os.Exit(0)
	} else if err != nil {
		fmt.Println(c.Name+":", err)
		fmt.Println()
		c.Flags.PrintUsage()
		os.Exit(0)
	}
}

/// Cmds is a map which contains defined utilities
var Cmds map[string]Commander

func init() {
	// Initialize Cmds
	Cmds = make(map[string]Commander)

	// Declare implemented utilities
	Cmds["mkdir"] = new(MKDIR)
	Cmds["rm"] = new(RM)
	//Cmds["touch"] = new(TOUCH)
}
