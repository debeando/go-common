package os

import (
	"os"
	"path/filepath"
)

func Args() []string {
	return os.Args[1:]
}

func ExecutableName() string {
	e, _ := os.Executable()
	return filepath.Base(e)
}
