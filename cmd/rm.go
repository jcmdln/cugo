package cmd

import (
	"fmt"
	"os"

	"github.com/markedhero/flagit"
)

type RM struct {
	Commander          // Parent interface
	Data      []string // Remaining data in args after flagit.Parse()

	// Local Flags
	Name        string
	About       string
	Interactive bool
	Verbose     bool
}

func (c *RM) Init() {
	var (
		Name = "rm"
		//About = "Creates the specified directories if they do not already exist"

		Force       = false
		Interactive = false
		Recursive   = false
		Verbose     = false

		err error
	)

	// Start new instance of and declare flags
	flags := flagit.NewFlag()
	flags.Bool(&Force, []string{"-f", "--force"},
		"Skip prompts and ignore warnings")
	flags.Bool(&Interactive, []string{"-i", "--interactive"},
		"Prompt before performing each action")
	flags.Bool(&Recursive, []string{"-r", "-R", "--recursive"},
		"Remove directories and their contents recursively")
	flags.Bool(&Verbose, []string{"-v", "--verbose"},
		"Print a message when actions are taken")

	c.Data, err = flags.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("cugo: mkdir:", err)
	}

	if len(os.Args) <= 2 {
		fmt.Println("cugo:", Name+":", "missing operand")
		flags.PrintUsage()
		os.Exit(0)
	}
}

func (c RM) Main() int {
	for _, target := range c.Data {
		if _, err := os.Stat(target); err == nil {
			if c.Interactive == true {
			}
			//os.RemoveAll(target)
			if c.Verbose == true {
				fmt.Println("Deleting", target)
			}
		}
	}

	return 0
}
