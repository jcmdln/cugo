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
	Use   string // Usage example
	About string // Description
	Man   string // Manual

	Flags *flagit.Flag
	Data  []string
}

func (c *Command) Args(opt int) {
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

var Cmds map[string]Commander

func init() {
	Cmds = make(map[string]Commander)
	Cmds["mkdir"] = new(MKDIR)
	Cmds["rm"] = new(RM)
}
