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
	_, h := s.Size()
	EmitStrMid(h-1, tcell.StyleDefault, currentDir)

	contents, err := os.ReadDir(currentDir)
	if err != nil { fmt.Println("Error reading directory:", err); return}

	if !displayHiddenFiles { contents = filterHiddenContents(contents) }

	maxFilenameLen := printContents(contents)

	printSelectionKeyHints(maxFilenameLen, contents)

	s.Show()
}

func printContents(contents []os.DirEntry)(int) {
	_, h := s.Size()

	skipOffset := -1
	maxFilenameLen := 0
	for i, file := range contents {
		if i % 5 == 0 { skipOffset += 1 }
		if file.IsDir() {
			EmitStrMid(i + skipOffset, tcell.StyleDefault, file.Name() + "/")
		} else {
			EmitStrMid(i + skipOffset, tcell.StyleDefault, file.Name())
		}

		if len(file.Name()) > maxFilenameLen { maxFilenameLen = len(file.Name()) }
		if i >= h-5 { break }
	}

	return maxFilenameLen
}

func printSelectionKeyHints(maxFilenameLen int, contents []os.DirEntry) {
	w, h := s.Size()

	skipOffset := -1
	if displayHints {
		keys := "asdfghjkl;zxcvbnm,./"
		for i := 0; i < len(contents); i++ {
			if i % 5 == 0 { skipOffset += 1 }
			EmitStr(w/2 - maxFilenameLen/2 - 3, i + skipOffset, tcell.StyleDefault, string(keys[i]) + ")")
			if i >= h-5 || i >= 19 { break }
		}
	}

}
	
func filterHiddenContents(contents []os.DirEntry) []os.DirEntry {
    filteredFiles := []os.DirEntry{}

    for _, file := range contents {
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

	contents, err := os.ReadDir(currentDir)
	if err != nil { fmt.Println("Error reading directory:", err); return}

	if !displayHiddenFiles { contents = filterHiddenContents(contents) }

	if(index < len(contents) && contents[index].IsDir()) {
		err = os.Chdir(contents[index].Name())
		if err != nil { fmt.Println("Error changing directory:", err); return}

		PrintCurrentDir()
	}
}
