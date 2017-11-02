package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/markedhero/flagit"
)

type MKDIR struct {
	Commander

	Name  string
	About string
	Data  []string // Remaining data in args after flagit.Parse()

	// Local Flags
	Mode    int
	Parents bool
	Verbose bool
}

func (c *MKDIR) Init() {
	Name, About := "mkdir", "Creates the specified directories if they do not already exist"

	Mode, Parents, Verbose := 0777, false, false
	flags := flagit.NewFlag()
	flags.Int(&Mode, []string{"-m", "--mode"},
		"Set folder permissions to specified MODE value")
	flags.Bool(&Parents, []string{"-p", "--parents"},
		"Create missing parent directories")
	flags.Bool(&Verbose, []string{"-v", "--verbose"},
		"Print a message when actions are taken")

	c.Data, _ = flags.Parse(os.Args[2:])

	if len(os.Args) <= 2 {
		fmt.Println("cugo:", Name+":", "missing operand")
		fmt.Println(Name, "-", About)
		flags.PrintUsage()
		os.Exit(0)
	}
}

func (c MKDIR) Main() int {
	for _, dir := range c.Data {
		_, err := os.Stat(dir)
		exists := os.IsExist(err)
		if err == nil {
			fmt.Println("cugo:", c.Name+":", err)
			return 0
		} else if exists {
			fmt.Println("cugo:", c.Name+":", "directory exists")
			return 0
		}

		if c.Parents == false {
			if !strings.Contains(dir, "/") {
				os.Mkdir(dir, os.FileMode(c.Mode))
				if c.Verbose == true {
					fmt.Println("Created directory", "'"+dir+"'",
						"with permissions", os.FileMode(uint32(c.Mode)))
				}
			} else {
				fmt.Println("cugo:", c.Name+":", err)
			}
		} else {
			os.MkdirAll(dir, os.FileMode(c.Mode))
			if c.Verbose == true {
				fmt.Println("Created directory", "'"+dir+"'",
					"with permissions", os.FileMode(uint32(c.Mode)))
			}
		}
	}

	return 0
}
