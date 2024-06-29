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

	if !displayHiddenFiles { files = filterHiddenFiles(files) }

	w, h := s.Size()
	EmitStrMid(h-1, tcell.StyleDefault, currentDir)
	skipOffset := -1
	maxFilenameLen := 0
	for i, file := range files {
		if i % 5 == 0 { skipOffset += 1 }
		if file.IsDir() {
			EmitStrMid(i + skipOffset, tcell.StyleDefault, file.Name() + "/")
		} else {
			EmitStrMid(i + skipOffset, tcell.StyleDefault, file.Name())
		}

		if len(file.Name()) > maxFilenameLen { maxFilenameLen = len(file.Name()) }
		if i >= h-5 { break }
	}

	skipOffset = -1
	if displayHints {
		keys := "asdfghjkl;zxcvbnm,./"
		for i := 0; i < len(files); i++ {
			if i % 5 == 0 { skipOffset += 1 }
			EmitStr(w/2 - maxFilenameLen/2 - 3, i + skipOffset, tcell.StyleDefault, string(keys[i]) + ")")
			if i >= h-5 || i >= 19 { break }
		}
	}

	s.Show()
}

	
func filterHiddenFiles(files []os.DirEntry) []os.DirEntry {
    filteredFiles := []os.DirEntry{}

    for _, file := range files {
        if file.Name()[0] != '.' {
            filteredFiles = append(filteredFiles, file)
        }
    }

    return filteredFiles
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

	if !displayHiddenFiles { files = filterHiddenFiles(files) }

	if(index < len(files) && files[index].IsDir()) {
		err = os.Chdir(files[index].Name())
		if err != nil { fmt.Println("Error changing directory:", err); return}

		PrintCurrentDir()
	}
}
