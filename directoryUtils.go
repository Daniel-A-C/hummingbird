package main

import (
	"fmt"
	"os"
    "path/filepath"
	"github.com/gdamore/tcell/v2"
)

func PrintCurrentDir() {
	s.Clear()

	currentDir, err := os.Getwd()
	if err != nil { fmt.Println("Error reading directory:", err); return}

	files, err := os.ReadDir(currentDir)
	if err != nil { fmt.Println("Error reading directory:", err); return}

	_, h := s.Size()
	EmitStrMid(h-1, tcell.StyleDefault, currentDir)
	skipOffset := -1
	for i, file := range files {
		if i % 5 == 0 { skipOffset += 1 }
		if file.IsDir() {
			EmitStrMid(i + skipOffset, tcell.StyleDefault, file.Name() + "/")
		} else {
			EmitStrMid(i + skipOffset, tcell.StyleDefault, file.Name())
		}
	}
	s.Show()
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
