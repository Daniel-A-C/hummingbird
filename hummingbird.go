package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func RunHummingbird(s tcell.Screen) (string){ 
	run := true
	w, _ := s.Size()

	EmitStr(s, w/2 - 5, 2, tcell.StyleDefault, "hummingbird")

	currentDir, err := os.Getwd()
	if err != nil { EmitStr(s, 0, 0, tcell.StyleDefault, "wd error"); return "x"}

	EmitStr(s, w/2 - 5, 3, tcell.StyleDefault, currentDir)

	s.Show()

	for run {
		if s.HasPendingEvent() {
			switch ev := s.PollEvent().(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
					run = false
				}
			}
		}

		time.Sleep(1 * time.Millisecond)
	}

	return "/Users/danielcarroll/code/hummingbird/testDir"
}

func main() {
	s := InitScreen()

	result := RunHummingbird(s)
	fmt.Print(result)

	s.Fini()

	//fmt.Print("x")
	//fmt.Print("/Users/danielcarroll/code/hummingbird/testDir")

/*
	files, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Contents of", currentDir)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("[DIR]", file.Name())
		} else {
			fmt.Println(file.Name())
		}
	}

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
