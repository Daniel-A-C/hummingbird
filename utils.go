package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

type Timer struct {
	startTime time.Time
}

// start/reset function
func (t *Timer) StartTimer() {
	t.startTime = time.Now()
}

// elapsedTime function
func (t *Timer) ElapsedTime() time.Duration {
	return time.Since(t.startTime)
}

func InitScreen() tcell.Screen { // Return type tcell.Screen, not (sc tcell.Screen)
	screen, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	screen.Clear()
	return screen
}

func EmitStrMid(y int, style tcell.Style, str string) {
	w, _ := s.Size()
	// Ensure string doesn't overflow; truncate if necessary, or handle differently
	// For simplicity, this basic centering is kept.
	strLen := runewidth.StringWidth(str) // Use runewidth for accurate length
	EmitStr(w/2-strLen/2, y, style, str)
}

func EmitStr(x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' ' // Assign placeholder if zero width
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

// PrintSelectionKeyHints displays the key mapping hints next to directory contents.
// maxFilenameLen: The length of the longest filename displayed, for alignment.
// numItemsDisplayed: The actual number of items shown on screen by printContents.
func PrintSelectionKeyHints(maxFilenameLen, numItemsDisplayed int) {
	if !displayHints || numItemsDisplayed == 0 {
		return
	}

	w, h := s.Size()
	if h == 0 {
		return
	}

	keys := "asdfghjkl;zxcvbnm,./"
	skipOffset := -1

	for i := 0; i < numItemsDisplayed && i < len(keys); i++ {
		if i%5 == 0 {
			skipOffset += 1
		}

		yPos := i + skipOffset
		// Ensure hints don't draw over the bottom reserved lines
		if yPos >= h-2 {
			break
		}

		hintStr := string(keys[i]) + ")"
		// Calculate x position for the hint: to the left of the centered content block
		// Center of screen: w/2
		// Start of centered block of filenames: w/2 - maxFilenameLen/2
		// Position hint 3 chars to the left of that: w/2 - maxFilenameLen/2 - 3
		// (2 for "X)" and 1 for space)
		hintX := w/2 - maxFilenameLen/2 - (runewidth.StringWidth(hintStr) + 1) // +1 for space
		if hintX < 0 {                                                         // Prevent negative x
			hintX = 0
		}

		EmitStr(hintX, yPos, tcell.StyleDefault.Dim(true), hintStr) // Dim style for hints
	}
}
