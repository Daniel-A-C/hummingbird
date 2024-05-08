package main

import (
	"fmt"
	"os"
	//"time"
    "path/filepath"
	"github.com/gdamore/tcell/v2"
)

var displayHiddenFiles = false
var s tcell.Screen

func PrintCurrentDir() {
	s.Clear()

	currentDir, err := os.Getwd()
	if err != nil { fmt.Println("Error reading directory:", err); return}

	files, err := os.ReadDir(currentDir)
	if err != nil { fmt.Println("Error reading directory:", err); return}

	_, h := s.Size()
	EmitStrMid(s, h-1, tcell.StyleDefault, currentDir)
	skipOffset := -1
	for i, file := range files {
		if i % 5 == 0 { skipOffset += 1 }
		if file.IsDir() {
			EmitStrMid(s, i + skipOffset, tcell.StyleDefault, file.Name() + "/")
		} else {
			EmitStrMid(s, i + skipOffset, tcell.StyleDefault, file.Name())
		}
	}
	s.Show()
}

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

func GoUpDir() {

	currentDir, err := os.Getwd()
	if err != nil { fmt.Println("Error reading directory:", err); return}

	err = os.Chdir(filepath.Dir(currentDir))
	if err != nil { fmt.Println("Error reading directory:", err); return}
	PrintCurrentDir()
}

func ChangeDir(index int) {

	currentDir, err := os.Getwd()
	if err != nil { fmt.Println("Error reading directory:", err); return}

	files, err := os.ReadDir(currentDir)
	if err != nil { fmt.Println("Error reading directory:", err); return}

	if(index < len(files) && files[index].IsDir()) {
		err = os.Chdir(files[index].Name())
		if err != nil { fmt.Println("Error changing directory:", err); return}

		PrintCurrentDir()
	}

}

func main() {
	s = InitScreen()

	result := RunHummingbird()
	fmt.Print(result)

	s.Fini()
}
