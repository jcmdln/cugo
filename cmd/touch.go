package cmd

import (
	"github.com/markedhero/flagit"
)

type TOUCH struct {
	Command

	AccessTime   bool
	Create       bool
	DateTime     string
	ModifiedTime bool
	Time         int
}

func (c *TOUCH) Init(flag []string) {
	c.Name = "touch"
	c.About = ""
	c.Use = "[-acm] [-r ref_file|-t TIME|-d date_time] FILES..."

	c.AccessTime = false
	c.Create = true
	c.ModifiedTime = false
	c.DateTime = ""
	c.Time = 0

	c.Flags = flagit.NewFlag()
	c.Flags.Bool(&c.AccessTime, []string{"-a"},
		"Change the access time of a file if -m is also specified")
	c.Flags.Bool(&c.Create, []string{"-c"},
		"Do not create the specified file if it doesn't exist")
	c.Flags.String(&c.DateTime, []string{"-d"},
		"")
	c.Flags.Bool(&c.ModifiedTime, []string{"m"},
		"Changed the modification time of a file if -a is also specified")
}

func (c *TOUCH) Main() int {
	return 0
}
