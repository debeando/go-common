package terminal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"golang.org/x/sys/unix"
)

const (
	// Reset all custom styles.
	RESET = "\033[0m"
	// Reset to default color.
	RESET_COLOR = "\033[39;49m"
	// Return cursor to start of line and clean it.
	LINE_RESET = "\r\033[K"
	// Erase the entire line.
	LINE_ERASE = "\x1B[2K"
	// Restore screen.
	SCREEN_RESTORE = "\x1B[?47l"
	// Save screen.
	SCREEN_SAVE = "\x1B[?47h"
	// Clear complete screen.
	SCREEN_CLEAR = "\033[2J"
	// Return cursor to top left.
	CURSOR_RESET = "\033[H"
	// Hide cursor.
	CURSOR_HIDE = "\x1b[?25l"
	// Show cursor.
	CURSOR_SHOW = "\x1b[?25h"
	// Erase from cursor to beginning of screen.
	CURSOR_ERASE_FBS = "\x1B[1J"
	// Erase from cursor to end of screen.
	CURSOR_ERASE_FES = "\x1B[0J"
	// Set new cursor position on screen.
	CURSOR_SET_POSITION = "\x1B[%d;%dH"
	// Save the cursor position.
	CURSOR_SAVE_POSITION = "\x1B7"
	// Restore the cursor position.
	CURSOR_RESTORE_POSITION = "\x1B8"
)

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)
var Screen *bytes.Buffer = new(bytes.Buffer)

func Clear() {
	// Clear screen
	Output.WriteString(SCREEN_CLEAR)
	// Return cursor to top left
	Output.WriteString(CURSOR_RESET)
}

func Reset() {
	Output.WriteString(RESET)
	Output.WriteString(RESET_COLOR)
	Output.WriteString(LINE_RESET)
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

func ScreenSave() {
	Output.WriteString(SCREEN_SAVE)
}

func ScreenRestore() {
	Output.WriteString(SCREEN_RESTORE)
}

func LineReset() {
	Output.WriteString(LINE_RESET)
}

func LineErase() {
	Output.WriteString(LINE_ERASE)
}

func CursorHide() {
	Output.WriteString(CURSOR_HIDE)
}

func CursorShow() {
	Output.WriteString(CURSOR_SHOW)
	Flush()
}

func CursorRestore() {
	Output.WriteString(CURSOR_RESTORE_POSITION)
	Flush()
}

func CursorSave() {
	Output.WriteString(CURSOR_SAVE_POSITION)
}

func CursorSet(x, y int) {
	Output.WriteString(fmt.Sprintf(CURSOR_SET_POSITION, x, y))
}

func CursorEraseToEndScreen() {
	Output.WriteString(CURSOR_ERASE_FES)
}

func CursorEraseToBeginningScreen() {
	Output.WriteString(CURSOR_ERASE_FBS)
}

func Refresh(wait int, f func()) {
	Reset()
	Clear()
	CursorHide()
	Flush()

	for {
		CursorSave()
		LineErase()
		CursorEraseToEndScreen()
		ScreenSave()
		CursorEraseToBeginningScreen()
		ScreenRestore()
		defer CursorRestore()

		CursorSet(0, 0)
		Flush()
		f()
		time.Sleep(time.Duration(wait) * time.Second)
	}

	CursorShow()
}

func getWinsize() (*unix.Winsize, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
