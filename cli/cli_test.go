package cli_test

import (
	"testing"

	"github.com/debeando/go-common/cli"
	// "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Log("demo")

	o := cli.Load()
	o.Args.Add(cli.Arg{Name: "foo", Description: "Foo method for test."})
	o.Args.Add(cli.Arg{Name: "bar", Description: "Bar method for test.", Required: true})
	o.Args.Add(cli.Arg{Name: "baz", Description: "Baz method for test."})

	t.Log(o.Help())
}
