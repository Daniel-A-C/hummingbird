package main

import (
	"fmt"
	"os"
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

type Timer struct {
    startTime time.Time 
}

// start/reset function
func (t *Timer) StartTimer() {
    t.startTime = time.Now()
}

// elapsedTime function
func (t *Timer) ElapsedTime() time.Duration {
    return time.Since(t.startTime)  
}

func InitScreen() (sc tcell.Screen) {
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	s.Clear()

	return s
}


func EmitStrMid(s tcell.Screen, y int, style tcell.Style, str string) {
	w, _ := s.Size()
	EmitStr(s, w/2 - len(str)/2, y, style, str)
}

func EmitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

