// Package term provides ansi escape sequence helpers.
package term

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/buger/goterm"
	"github.com/mattn/go-isatty"
)

// Renderer returns a render function which computes
// the diff of each update required line-by-line,
// and performs the changes.
//
// This provides a nicer user experience than
// simply running re-rendering with ClearAll().
func Renderer() func(string) {
	var prev string
	return func(curr string) {
		for i, line := range linesChanged(curr, prev) {
			if line != "" {
				MoveTo(i+1, 1)
				ClearLineEnd()
				fmt.Printf("%s", line)
			}
		}
		prev = curr
	}
}

// linesChanged returns the lines changed, while unchanged
// lines are simply empty strings.
func linesChanged(curr, prev string) (lines []string) {
	currLines := strings.Split(curr, "\n")
	prevLines := strings.Split(prev, "\n")

	for i, line := range currLines {
		if len(prevLines) > i && line == prevLines[i] {
			lines = append(lines, "")
		} else {
			lines = append(lines, line)
		}
	}

	return
}

// strip regexp.
var strip = regexp.MustCompile(`\x1b\[(\d+[;m])+`)

// Strip ansi escape sequences.
func Strip(s string) string {
	return strip.ReplaceAllString(s, "")
}

// Length of characters with ansi escape sequences stripped.
func Length(s string) (n int) {
	for range Strip(s) {
		n++
	}
	return
}

// CenterLine a line of text.
func CenterLine(s string) string {
	r := strings.Repeat
	w, h := Size()
	size := Length(s)
	xpad := int(math.Abs(float64((w - size) / 2)))
	ypad := h / 2
	return r("\n", ypad) + r(" ", xpad) + s + r("\n", ypad)
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

// IsTerminal returns true if stdout isatty.
func IsTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}
