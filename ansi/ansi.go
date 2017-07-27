// Package ansi provies ansi escape sequence helpers.
package ansi

import (
	"fmt"
	"regexp"

	"github.com/buger/goterm"
)

// strip regexp.
var strip = regexp.MustCompile(`\x1b\[(\d+[;m])+`)

// Strip ansi escape sequences.
func Strip(s string) string {
	return strip.ReplaceAllString(s, "")
}

// Size returns the width and height.
func Size() (w int, h int) {
	w = goterm.Width()
	h = goterm.Height()
	return
}

// ClearAll clears the screen.
func ClearAll() {
	fmt.Printf("\033[2J")
	MoveTo(1, 1)
}

// ClearLine clears the entire line.
func ClearLine() {
	fmt.Printf("\033[2K")
}

// ClearLineEnd clears to the end of the line.
func ClearLineEnd() {
	fmt.Printf("\033[0K")
}

// ClearLineStart clears to the start of the line.
func ClearLineStart() {
	fmt.Printf("\033[1K")
}

// MoveTo moves the cursor to (x, y).
func MoveTo(x, y int) {
	fmt.Printf("\033[%d;%df", x, y)
}

// HideCursor hides the cursor.
func HideCursor() {
	fmt.Printf("\033[?25l")
}

// ShowCursor shows the cursor.
func ShowCursor() {
	fmt.Printf("\033[?25h")
}
