package main

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell/v2"
)

var s tcell.Screen

// Eventually should belong in some sort of "settings" feature.
var displayHiddenFiles = false

func RunHummingbird() (string){ 
    inputMap := map[byte]int{
        'a': 0, 's': 1, 'd': 2, 'f': 3, 'g': 4, 'h': 5, 'j': 6, 'k': 7, 'l': 8, ';': 9,
        'z': 10, 'x': 11, 'c': 12, 'v': 13, 'b': 14, 'n': 15, 'm': 16, ',': 17, '.': 18, '/': 19,
    }

	run := true
	PrintCurrentDir()

	for run {
		if s.HasPendingEvent() {
			switch ev := s.PollEvent().(type) {
			case *tcell.EventKey:
				index, ok := inputMap[byte(ev.Rune())] // If the input is in the map
				if ok {
					ChangeDir(index)
				} else if string(ev.Rune()) == "e" {
					GoUpDir()
				} else if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
					run = false
				} else if string(ev.Rune()) == "y" {

				}
			}
		}
	}

	currentDir, err := os.Getwd()
	if err != nil { fmt.Println("Error reading directory:", err); return "x"}
	return currentDir
}


func main() {
	s = InitScreen()

	result := RunHummingbird()
	fmt.Print(result)

	s.Fini()
}
