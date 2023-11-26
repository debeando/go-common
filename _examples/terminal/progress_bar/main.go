package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/debeando/go-common/terminal"
)

func main() {
	terminal.Reset()
	terminal.Clear()
	terminal.Flush()

	for i := 0; i < 10; i++ {
		// fmt.Print("\x1B7")       // Save the cursor position
		// fmt.Print("\x1B[2K")     // Erase the entire line
		// fmt.Print("\x1B[0J")     // Erase from cursor to end of screen
		// fmt.Print("\x1B[?47h")   // Save screen
		// fmt.Print("\x1B[1J")     // Erase from cursor to beginning of screen
		// fmt.Print("\x1B[?47l")   // Restore screen
		// defer fmt.Print("\x1B8") // Restore the cursor position util new size is calculated

		terminal.Cursor(0, i)
		fmt.Println(strings.Repeat("=", i))
		// fmt.Printf("Progress: [\x1B[33m%3d%%\x1B[0m] %s", 2, "===")
		time.Sleep(100 * time.Millisecond)
	}
}
