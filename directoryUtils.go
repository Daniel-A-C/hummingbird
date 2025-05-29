package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
	"path/filepath" // Needed for filepath.Join
)

// PrintCurrentDir clears the screen and displays the current directory contents.
func PrintCurrentDir() {
	s.Clear()

	currentDir, err := os.Getwd()
	if err != nil {
		errorMsg := fmt.Sprintf("Error reading directory: %v", err)
		_, h := s.Size()
		if h > 0 {
			EmitStrMid(0, tcell.StyleDefault.Foreground(tcell.ColorRed), errorMsg)
		}
		s.Show()
		return
	}

	contents, err := os.ReadDir(currentDir)
	if err != nil {
		errorMsg := fmt.Sprintf("Error reading dir contents: %v", err)
		_, h := s.Size()
		if h > 0 {
			EmitStrMid(1, tcell.StyleDefault.Foreground(tcell.ColorRed), errorMsg) // Line 1 for this error
		}
		s.Show()
		return
	}

	if !displayHiddenFiles {
		contents = filterHiddenContents(contents)
	}

	var maxFilenameLen, numItemsDisplayed int
	if len(contents) > 0 {
		maxFilenameLen, numItemsDisplayed = printContents(contents)
	} else {
		// No contents to display, perhaps show an (empty) message
		// EmitStrMid(0, tcell.StyleDefault, "(empty)")
		maxFilenameLen = 0
		numItemsDisplayed = 0
	}

	PrintSelectionKeyHints(maxFilenameLen, numItemsDisplayed)

	// Display current directory path at the bottom
	_, screenH := s.Size()
	if screenH > 0 {
		EmitStrMid(screenH-1, tcell.StyleDefault, currentDir)
	}

	s.Show()
}

// printContents displays the directory entries.
// Returns:
//   - maxFilenameLen: The maximum length of the filenames/dirnames displayed.
//   - itemsPrinted: The number of items actually printed to the screen.
func printContents(contents []os.DirEntry) (int, int) {
	_, h := s.Size()
	if h == 0 {
		return 0, 0
	} // No space to print

	skipOffset := -1
	maxFilenameLen := 0
	itemsPrinted := 0

	for i, file := range contents {
		if i%5 == 0 {
			skipOffset += 1
		}

		yPos := i + skipOffset
		// Reserve last two lines (h-1 for path, h-2 for margin/other status)
		if yPos >= h-2 {
			break
		}

		name := file.Name()
		if file.IsDir() {
			EmitStrMid(yPos, tcell.StyleDefault.Bold(true), name+"/")
		} else {
			EmitStrMid(yPos, tcell.StyleDefault, name)
		}

		if len(name) > maxFilenameLen {
			maxFilenameLen = len(name)
		}
		itemsPrinted++
	}
	return maxFilenameLen, itemsPrinted
}

// filterHiddenContents removes entries starting with '.' from the list.
func filterHiddenContents(contents []os.DirEntry) []os.DirEntry {
	filteredFiles := make([]os.DirEntry, 0, len(contents))
	for _, file := range contents {
		if file.Name()[0] != '.' {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles
}

// GoUpDir changes the current directory to its parent.
func GoUpDir() {
	currentDir, err := os.Getwd()
	if err != nil {
		// Consider displaying error in TUI status line
		return
	}
	parentDir := filepath.Dir(currentDir)
	if parentDir == currentDir { // Already at root
		return
	}
	err = os.Chdir(parentDir)
	if err != nil {
		// Consider displaying error in TUI status line
		return
	}
	PrintCurrentDir()
}

// ChangeDir attempts to change into a directory or select a file based on index.
// Returns true if a file was selected (and program should exit), false otherwise.
func ChangeDir(index int) bool {
	currentDir, err := os.Getwd()
	if err != nil {
		return false
	}

	contents, err := os.ReadDir(currentDir)
	if err != nil {
		return false
	}

	if !displayHiddenFiles {
		contents = filterHiddenContents(contents)
	}

	if index < 0 || index >= len(contents) {
		return false // Index out of bounds
	}

	selectedEntry := contents[index]
	targetPath := filepath.Join(currentDir, selectedEntry.Name())

	if selectedEntry.IsDir() {
		err = os.Chdir(targetPath)
		if err != nil {
			return false
		}
		PrintCurrentDir()
		return false // Directory changed, continue running
	} else { // It's a file
		selectedFileToExitWith = targetPath
		return true // File selected, signal to exit
	}
}
