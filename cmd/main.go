package cmd

import (
	"fmt"
)

type Commander interface {
	Init()     // Initialization
	Main() int // Execution
}

type Command struct {
	Name string
	Use  string
}

func (c Command) Args(index int, numFlags int) bool {
	if numFlags < index {
		fmt.Println(c.Name+":", "Not enough operands")
		fmt.Println("Usage:", c.Name, c.Use)
		return false
	}
	return true
}

var Cmds map[string]Commander

func init() {
	Cmds = make(map[string]Commander)
	Cmds["mkdir"] = new(MKDIR)
}
