package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
)

var s tcell.Screen

var displayHints = true
var displayHiddenFiles = false
var selectedFileToExitWith string = "" // Stores the path of a selected file for exit

func main() {
	var err error
	s, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create screen: %v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize screen: %v\n", err)
		os.Exit(1)
	}
	// Assuming InitScreen was a placeholder for tcell.NewScreen and s.Init()
	// If InitScreen() had other critical setup, that needs to be preserved/merged.

	result := runHummingbird()
	s.Fini() // Finalize screen AFTER runHummingbird completes

	// Print the result to stdout AFTER s.Fini() so terminal is back to normal
	fmt.Print(result)
}

func runHummingbird() string {
	run := true
	selectedFileToExitWith = "" // Reset at the beginning of a run

	PrintCurrentDir() // Initial display

	for run {
		ev := s.PollEvent() // Blocks until an event occurs
		if ev == nil {
			// This can happen if the screen is finalized concurrently or an error occurs
			break
		}

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()          // Synchronize screen size
			PrintCurrentDir() // Redraw content
		case *tcell.EventKey:
			// Exit keys
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC ||
				(ev.Key() == tcell.KeyRune && ev.Rune() == 'q') {
				run = false
				// selectedFileToExitWith will be empty, so current directory will be returned
				break // Break from switch, loop condition 'run' will handle exit
			}
			// Pass the event to respondToKeyPress. If it returns false, we should stop.
			if !respondToKeyPress(ev) {
				run = false // A file was selected, or another reason to stop from respondToKeyPress
			}
		}
	}

	// After loop finishes, decide what to return
	if selectedFileToExitWith != "" {
		return selectedFileToExitWith // A file was selected
	}

	// Default: return current directory on normal quit (Esc, q, Ctrl+C)
	currentDir, err := os.Getwd()
	if err != nil {
		// This error will print to stderr after Fini if Fini is called before print in main
		// Or print to a potentially messed up terminal if printed directly here before Fini
		// For now, let's assume errors should go to stderr.
		fmt.Fprintf(os.Stderr, "Error reading current directory at exit: %v\n", err)
		return "error_getting_wd_at_exit" // Indicate an error
	}
	return currentDir
}

// respondToKeyPress handles key presses for navigation and actions.
// It returns true to continue running, false to stop the main loop (e.g., file selected).
func respondToKeyPress(ev *tcell.EventKey) bool {
	keyChar := string(ev.Rune()) // Get the character for rune-based key presses

	// Map for directory/file selection keys ('a', 's', 'd', etc.)
	var inputMap = map[string]int{
		"a": 0, "s": 1, "d": 2, "f": 3, "g": 4, "h": 5, "j": 6, "k": 7, "l": 8, ";": 9,
		"z": 10, "x": 11, "c": 12, "v": 13, "b": 14, "n": 15, "m": 16, ",": 17, ".": 18, "/": 19,
	}

	index, isMappedKey := inputMap[keyChar]
	if isMappedKey {
		// ChangeDir now returns true if a file was selected (and global path is set), false otherwise.
		if ChangeDir(index) { // ChangeDir will set selectedFileToExitWith if a file is chosen
			return false // Signal to stop running, as a file was selected.
		}
		// If ChangeDir returned false, it means a directory was navigated or an invalid index was chosen.
		// PrintCurrentDir is called by ChangeDir on successful directory change.
		// No need to call PrintCurrentDir here unless ChangeDir fails silently.
		return true // Continue running.
	}

	// Handle other specific character keys
	switch keyChar {
	case "e":
		GoUpDir()
	case "y":
		displayHiddenFiles = !displayHiddenFiles
		PrintCurrentDir()
	case "u":
		displayHints = !displayHints
		PrintCurrentDir()
	case "?":
		RunSettingsMenu()
		PrintCurrentDir() // Redraw main UI after settings menu closes
	default:
		// Key is not 'q'/Esc/Ctrl+C (handled in main loop),
		// not in inputMap, and not e, y, u, ?.
		// It's an unhandled key press. Do nothing.
	}
	return true // Default: continue running
}
