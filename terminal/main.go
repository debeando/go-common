package terminal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"

	"golang.org/x/sys/unix"
)

const (
	// Clear screen
	CLEAR = "\033[2J"
	// Return cursor to top left
	RESET_CURSOR = "\033[H"
	// Reset all custom styles
	RESET = "\033[0m"
	// Reset to default color
	RESET_COLOR = "\033[39;49m"
	// Return cursor to start of line and clean it
	RESET_LINE = "\r\033[K"
)

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)
var Screen *bytes.Buffer = new(bytes.Buffer)

func Clear() {
	// Clear screen
	Output.WriteString(CLEAR)
	// Return cursor to top left
	Output.WriteString(RESET_CURSOR)
}

func Reset() {
	fmt.Print(RESET)
	fmt.Print(RESET_COLOR)
	fmt.Print(RESET_LINE)
}

func Flush() {
	Output.Flush()
	Screen.Reset()
}

func Width() int {
	ws, err := getWinsize()
	if err != nil {
		return -1
	}

	return int(ws.Col)
}

func Height() int {
	ws, err := getWinsize()
	if err != nil {
		if errors.Is(err, unix.EOPNOTSUPP) {
			return math.MinInt32
		}
		return -1
	}
	return int(ws.Row)
}

func getWinsize() (*unix.Winsize, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
