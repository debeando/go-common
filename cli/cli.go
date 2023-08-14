package cli

import (
	"io"
	"flag"
	"fmt"
	"strings"
)

const (
	Type    uint8 = 0
	Boolean       = 1
	Integer       = 2
	String        = 3
)

type CLI struct {
	Description string
	Args        Args
	Usage       string
	Version     string
}

var options *CLI

func Load() *CLI {
	if options == nil {
		options = &CLI{}
	}
	return options
}

func (c *CLI) Help() string {
	var help string
	var required string
	var version string

	if len(c.Version) > 0 {
		version = fmt.Sprintf(" (%s)", c.Version)
	}

	help = fmt.Sprintf("%s%s\n", ExecutableName(), version)

	if len(c.Description) > 0 {
		help += fmt.Sprintf("  %s\n", c.Description)
	}

	if len(c.Usage) > 0 {
		help += "\nUsage:\n"
		help += fmt.Sprintf("  %s %s\n", ExecutableName(), c.Usage)
	}

	if len((*c).Args) > 0 {
		help += "\nOptions:\n"

		for _, arg := range c.Args {
			required = ""

			if arg.Required {
				required = " (Required)"
			}

			name   := arg.Name
			length := c.Args.NamesLength() - len(name)
			name   += strings.Repeat(" ", length)

			help += fmt.Sprintf("  --%s\t%s%s\n", name, arg.Description, required)
		}
	}

	return help
}

func (c *CLI) Lookup(name string) Arg {
	for _, arg := range c.Args {
		if arg.Name == name {
			return arg
		}
	}

	return Arg{}
}

func (c *CLI) Parser() {
	f := flag.NewFlagSet("default", flag.ContinueOnError)

	for _, arg := range c.Args {
		switch arg.Type {
		case Boolean:
			f.Bool(arg.Name, false, "")
		case Integer:
			f.Int(arg.Name, 0, "")
		case String:
			f.String(arg.Name, "", "")
		}
	}

	f.SetOutput(io.Discard)
	f.Parse(GetArgs())

	for index, arg := range c.Args {
		(*c).Args[index].Value = f.Lookup(arg.Name).Value.String()
	}
}
