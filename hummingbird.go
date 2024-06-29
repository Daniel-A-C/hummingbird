package main

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell/v2"
)

var s tcell.Screen

// Eventually should belong in some sort of "settings" feature.
var displayHiddenFiles = false
var displayHints = true

func main() {
	s = InitScreen()

	result := runHummingbird()
	fmt.Print(result)

	s.Fini()
}

func runHummingbird() (string){ 

	run := true
	PrintCurrentDir()

	for run {
		for !s.HasPendingEvent() { }

		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
				run = false
				break
			}
			respondToKeyPress(string(ev.Rune()))
		}
	}

	currentDir, err := os.Getwd()
	if err != nil { fmt.Println("Error reading directory:", err); return "x"}
	return currentDir
}

func respondToKeyPress(key string) {
	var inputMap = map[string]int{
		"a": 0, "s": 1, "d": 2, "f": 3, "g": 4, "h": 5, "j": 6, "k": 7, "l": 8, ";": 9,
		"z": 10, "x": 11, "c": 12, "v": 13, "b": 14, "n": 15, "m": 16, ",": 17, ".": 18, "/": 19,
	}

	index, ok := inputMap[key]
	if ok { // If the input is in the map
		ChangeDir(index)
	} else if key == "e" {
		GoUpDir()
	} else if key == "y" {
		displayHiddenFiles = !displayHiddenFiles
		PrintCurrentDir()
	} else if key == "u" { 
		displayHints = !displayHints
		PrintCurrentDir()
	}
}
