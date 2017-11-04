package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/markedhero/flagit"
)

type MKDIR struct {
	Command

	Help    bool
	Mode    int
	Parents bool
	Verbose bool
}

func (c *MKDIR) Init(flag []string) {
	c.Name = "mkdir"
	c.Use = "[OPTION]... DIRECTORY..."
	c.About = "Creates the specified directories if they do not already exist"

	c.Help = false
	c.Parents = false
	c.Verbose = false
	c.Mode = 0755

	flags := flagit.NewFlag()
	flags.Bool(&c.Help, []string{"-h", "--help"},
		"Prints this text")
	flags.Int(&c.Mode, []string{"-m", "--mode"},
		"Set folder permissions to specified MODE value")
	flags.Bool(&c.Parents, []string{"-p", "--parents"},
		"Create missing parent directories")
	flags.Bool(&c.Verbose, []string{"-v", "--verbose"},
		"Print a message when actions are taken")

	var err error
	c.Data, err = flags.Parse(os.Args[2:])
	a := func() {
		fmt.Println()
		flags.PrintUsage()
		os.Exit(0)
	}
	if !c.Args() || c.Help == true {
		fmt.Println(c.Name, "-", c.About)
		a()
	} else if err != nil {
		fmt.Println(c.Name+":", err)
		a()
	}
}

func (c MKDIR) Main() int {
	for _, dir := range c.Data {
		_, err := os.Stat(dir)
		exists := !os.IsNotExist(err)
		if exists {
			fmt.Println(c.Name+":", "'"+dir+"'", "already exists!")
			os.Exit(0)
		}

		v := func() {
			if c.Verbose == true {
				fmt.Println("Creating", "'"+dir+"'", "with permissions",
					os.FileMode(uint32(c.Mode)))
			}
		}

		if c.Parents == true {
			v()
			os.MkdirAll(dir, os.FileMode(c.Mode))
		} else if !c.Parents && !strings.Contains(dir, "/") {
			v()
			os.Mkdir(dir, os.FileMode(uint32(c.Mode)))
		} else if strings.Contains(dir, "/"); err != nil {
			fmt.Println(c.Name+":", err)
		}
	}

	return 0
}
