package cli

import (
	"os"
	"testing"

	"github.com/debeando/go-common/cli"
	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {
	os.Args = append(os.Args, "--foo=one")
	os.Args = append(os.Args, "--bar=two")
	os.Args = append(os.Args, "--baz")

	assert.Len(t, cli.GetArgs(), 7)
}

func TestExecutableName(t *testing.T) {
	assert.Equal(t, cli.ExecutableName(), "cli.test")
}
