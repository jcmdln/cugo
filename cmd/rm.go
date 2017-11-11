package cmd

import (
	"fmt"
	"os"

	"github.com/markedhero/flagit"
)

type RM struct {
	Command

	Force       bool
	Interactive bool
	Recursive   bool
	Verbose     bool
}

func (c *RM) Init(flag []string) {
	c.Name = "rm"
	c.Use = "[OPTION]... TARGETS..."
	c.About = "Creates the specified directories if they do not already exist"

	c.Force = false
	c.Interactive = false
	c.Recursive = false
	c.Verbose = false

	c.Flags = flagit.NewFlag()
	c.Flags.Bool(&c.Force, []string{"-f", "--force"},
		"Skip prompts and ignore warnings")
	c.Flags.Bool(&c.Interactive, []string{"-i", "--interactive"},
		"Prompt before performing each action")
	c.Flags.Bool(&c.Recursive, []string{"-r", "-R", "--recursive"},
		"Remove directories and their contents recursively")
	c.Flags.Bool(&c.Verbose, []string{"-v", "--verbose"},
		"Print a message when actions are taken")
	c.Args(2)
}

func (c RM) Main() int {
	for _, target := range c.Data {
		if _, err := os.Stat(target); err == nil {
			if c.Interactive == true {
				//
			}
			//os.RemoveAll(target)
			if c.Verbose == true {
				fmt.Println("Pretending to delete", target)
			}
		}
	}

	return 0
}
