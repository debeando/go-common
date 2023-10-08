package exec

import (
	"fmt"
	"os/exec"
	"syscall"

	"github.com/debeando/go-common/cast"
)

func PGrep(cmd string) int64 {
	stdout, _ := Command(fmt.Sprintf("/usr/bin/pgrep -f '%s'", cmd))

	return cast.StringToInt64(stdout)
}

func Command(cmd string) (stdout string, exitcode int) {
	out, err := exec.Command("/bin/bash", "-c", cmd).Output()

	if exitError, ok := err.(*exec.ExitError); ok {
		ws := exitError.Sys().(syscall.WaitStatus)
		exitcode = ws.ExitStatus()
	}

	stdout = string(out[:])
	return
}
