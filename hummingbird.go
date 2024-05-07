package main

import (
	"fmt"
	"os"
	"time"
    "path/filepath"
	"github.com/gdamore/tcell/v2"
)

func PrintCurrentDir(s tcell.Screen) {
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

func RunHummingbird(s tcell.Screen) (string){ 
    mapping := map[byte]int{
        'a': 0, 's': 1, 'd': 2, 'f': 3, 'g': 4, 'h': 5, 'j': 6, 'k': 7, 'l': 8, ';': 9,
        'z': 10, 'x': 11, 'c': 12, 'v': 13, 'b': 14, 'n': 15, 'm': 16, ',': 17, '.': 18, '/': 19,
    }

	run := true

	//EmitStr(s, 0, 0, tcell.StyleDefault, "hummingbird")
	PrintCurrentDir(s)

	for run {
		if s.HasPendingEvent() {
			switch ev := s.PollEvent().(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
					run = false
				}
				index, ok := mapping[byte(ev.Rune())]
				if ok {
					currentDir, err := os.Getwd()
					if err != nil { fmt.Println("Error reading directory:", err); return "x"}

					files, err := os.ReadDir(currentDir)
					if err != nil { fmt.Println("Error reading directory:", err); return "x"}

					if(index < len(files) && files[index].IsDir()) {
						err = os.Chdir(files[index].Name())
						if err != nil { fmt.Println("Error changing directory:", err); return "x"}

						PrintCurrentDir(s)
					}
					s.Show()
				}
				if string(ev.Rune()) == "e" {
					currentDir, err := os.Getwd()
					if err != nil { fmt.Println("Error reading directory:", err); return "x"}

					err = os.Chdir(filepath.Dir(currentDir))
					if err != nil { fmt.Println("Error reading directory:", err); return "x"}
					PrintCurrentDir(s)
				}
			}
		}

		time.Sleep(1 * time.Millisecond)
	}

	currentDir, err := os.Getwd()
	if err != nil { fmt.Println("Error reading directory:", err); return "x"}
	return currentDir
	//return "/Users/danielcarroll/code/hummingbird/testDir"
	//return "x"
}

func main() {
	s := InitScreen()

	result := RunHummingbird(s)
	fmt.Print(result)

	s.Fini()

/*

	fmt.Print("\nEnter the number of the directory to change to: ")
	var dirIndex int
	fmt.Scanln(&dirIndex)
	var targetDir string

	if dirIndex >= 1 && dirIndex <= len(files) {
		targetDir = files[dirIndex-1].Name() // Directories are listed first
		// ... change directory logic below
	} else {
		fmt.Println("Invalid selection.")
	}

	err = os.Chdir(targetDir)
	if err != nil {
		fmt.Println("Error changing directory:", err)
	} else {
		fmt.Println("Directory changed successfully.")
	}

	currentDir, err = os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
*/	
}
