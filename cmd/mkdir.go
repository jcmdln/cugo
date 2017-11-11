package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/markedhero/flagit"
)

type MKDIR struct {
	Command

	Mode    int
	Parents bool
	Verbose bool
}

func (c *MKDIR) Init(flag []string) {
	c.Name = "mkdir"
	c.Use = "[OPTION]... DIRECTORY..."
	c.About = "Creates the specified directories if they do not already exist"

	c.Mode = 0777
	c.Parents = false
	c.Verbose = false

	c.Flags = flagit.NewFlag()
	c.Flags.Int(&c.Mode, []string{"-m", "--mode"},
		"Set folder permissions to specified MODE value")
	c.Flags.Bool(&c.Parents, []string{"-p", "--parents"},
		"Create missing parent directories")
	c.Flags.Bool(&c.Verbose, []string{"-v", "--verbose"},
		"Print a message when actions are taken")
	c.Args(2)
}

func (c MKDIR) Main() int {
	for _, dir := range c.Data {
		_, err := os.Stat(dir)
		exists := !os.IsNotExist(err)
		if exists {
			fmt.Println(c.Name+":", "'"+dir+"'", "already exists!")
			return 0
		}

		if c.Parents == true {
			os.MkdirAll(dir, os.FileMode(c.Mode))
		} else if !c.Parents && !strings.Contains(dir, "/") {
			os.Mkdir(dir, os.FileMode(uint32(c.Mode)))
		} else if strings.Contains(dir, "/"); err != nil {
			fmt.Println(c.Name+":", err)
			return 0
		}

		if c.Verbose == true {
			fmt.Println("Created", "'"+dir+"'", "with permissions",
				os.FileMode(uint32(c.Mode)))
		}
	}

	return 0
}
