package cmd

import (
	"fmt"
	"os"
)

type Commander interface {
	Init(flag []string)
	Main() int
}

type Command struct {
	Data []string

	Name  string
	Use   string
	About string
	Man   string
}

func (c Command) Args() bool {
	if len(os.Args[2:]) < 1 {
		fmt.Println(c.Name+":", "Not enough operands")
		fmt.Println("Usage:", c.Name, c.Use)
		os.Exit(0)
	}
	return true
}

var Cmds map[string]Commander

func init() {
	Cmds = make(map[string]Commander)
	Cmds["mkdir"] = new(MKDIR)
}
