package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func main() {

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	//s.SetStyle(tcell.StyleDefault.
	//	Foreground(tcell.ColorBlack).
	//	Background(tcell.ColorWhite))
	s.Clear()

	snake := NewSnake(Coordinates{5, 5}, Coordinates{5, 6}, Coordinates{5, 7})
	//	w, _ := s.Size()

	st := tcell.StyleDefault

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				case tcell.KeyUp:
					snake.Turn(Up)
				case tcell.KeyDown:
					snake.Turn(Down)
				case tcell.KeyLeft:
					snake.Turn(Left)
				case tcell.KeyRight:
					snake.Turn(Right)
				}

			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 100):
			s.Clear()
			s.SetCell(5, 7, st, '*')

			headCoord := snake.Step()
			if headCoord.x == 5 && headCoord.y == 7 {
				snake.Eat(headCoord)
			} else {
				snake.Move(headCoord)
			}

			snake.Draw(func(c Coordinates) {
				s.SetCell(c.x, c.y, st, '#')
			})
		}

		s.Show()
	}

	s.Fini()
}
