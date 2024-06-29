package main

import (
	"github.com/gdamore/tcell/v2"
)

func RunSettingsMenu() {

	run := true
	printSettingsMenu()

	for run {
		for !s.HasPendingEvent() { }

		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || string(ev.Rune()) == "q" || ev.Key() == tcell.KeyCtrlC {
				run = false
				break
			}
			
		}
	}
	PrintCurrentDir()
}

func printSettingsMenu() {
	s.Clear()

	EmitStrMid(2, tcell.StyleDefault, "settings")

	s.Show()
}
