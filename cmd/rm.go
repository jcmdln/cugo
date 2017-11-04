package cmd

import (
	"fmt"
	"os"

	"github.com/markedhero/flagit"
)

type RM struct {
	Command
	Data []string

	Help        bool
	Force       bool
	Interactive bool
	Recursive   bool
	Verbose     bool
}

func (c *RM) Init() {
	c.Name = "rm"
	c.About = "Creates the specified directories if they do not already exist"

	var (
		Force       = false
		Help        = false
		Interactive = false
		Recursive   = false
		Verbose     = false

		err error
	)

	// Start new instance of and declare flags
	flags := flagit.NewFlag()
	flags.Bool(&Help, []string{"-h", "--help"},
		"Prints this text")
	flags.Bool(&Force, []string{"-f", "--force"},
		"Skip prompts and ignore warnings")
	flags.Bool(&Interactive, []string{"-i", "--interactive"},
		"Prompt before performing each action")
	flags.Bool(&Recursive, []string{"-r", "-R", "--recursive"},
		"Remove directories and their contents recursively")
	flags.Bool(&Verbose, []string{"-v", "--verbose"},
		"Print a message when actions are taken")

	c.Data, err = flags.Parse(os.Args[2:])
	a := func() {
		fmt.Println()
		flags.PrintUsage()
		os.Exit(0)
	}
	if !c.Args() || Help == true {
		fmt.Println(c.Name, "-", c.About)
		a()
	} else if err != nil {
		fmt.Println(c.Name+":", err)
		a()
	}
}

func (c RM) Main() int {
	for _, target := range c.Data {
		if _, err := os.Stat(target); err == nil {
			if c.Interactive == true {
			}
			//os.RemoveAll(target)
			if c.Verbose == true {
				fmt.Println("Pretending to delete", target)
			}
		}
	}

	return 0
}
