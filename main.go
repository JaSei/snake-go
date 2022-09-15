package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

type boxRoot struct {
	views.BoxLayout
}

type boxL struct {
	views.BoxLayout
}

var arena = &boxL{}
var info = &boxL{}
var root = &boxRoot{}
var app = &views.Application{}

func (m *boxRoot) HandleEvent(ev tcell.Event) bool {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			app.Quit()
			return true
		}
	}
	return m.BoxLayout.HandleEvent(ev)
}

func (m *boxL) HandleEvent(ev tcell.Event) bool {
	return false
}

func main() {
	title := views.NewText()
	//title.SetStyle(tcell.StyleDefault.
	//	Foreground(tcell.ColorBlack).
	//	Background(tcell.ColorYellow))
	title.SetText("Points: 10")
	//top := views.NewText()
	//mid := views.NewText()
	//bot := views.NewText()

	//top.SetText("Top-Right (0.0)\nLine Two")
	//mid.SetText("Center (0.7)")
	//bot.SetText("Bottom-Left (0.3)")
	//top.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).
	//	Background(tcell.ColorRed))
	//mid.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).
	//	Background(tcell.ColorLime))
	//bot.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).
	//	Background(tcell.ColorBlue))

	//top.SetAlignment(views.VAlignTop | views.HAlignRight)
	//mid.SetAlignment(views.VAlignCenter | views.HAlignCenter)
	//bot.SetAlignment(views.VAlignBottom | views.HAlignLeft)

	//v := views.
	root.SetOrientation(views.Horizontal)
	root.AddWidget(arena, 0)
	root.AddWidget(info, 0)
	info.AddWidget(title, 0)
	//info.AddWidget(top, 0)
	//info.AddWidget(mid, 0.7)
	//info.AddWidget(bot, 0.3)

	app.SetRootWidget(root)
	if e := app.Run(); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(1)
	}
	//	app.SetRootWidget(window)
	//	if e := app.Run(); e != nil {
	//		fmt.Fprintln(os.Stderr, e.Error())
	//		os.Exit(1)
	//	}
	//
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
				case tcell.KeyCtrlR:
					//init()
				}

			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	food := Coordinates{5, 7}

	end := false
loop:
	for {
		var dur time.Duration

		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond*100 - dur):
			if end {
				continue loop
			}

			startTime := time.Now()
			s.Clear()
			s.SetCell(food.x, food.y, st, '*')

			headCoord := snake.Step()

			if headCoord.x == 0 || headCoord.y == 0 {
				end = true
			}

			if snake.Containes(headCoord) {
				end = true
			}

			if end {
			} else if headCoord.x == food.x && headCoord.y == food.y {
				snake.Eat(headCoord)
				food.x++
				food.y++
			} else {
				snake.Move(headCoord)
			}

			snake.Draw(func(c Coordinates) {
				s.SetCell(c.x, c.y, st, '#')
			})

			if end {
				crash(s, headCoord)
			}

			s.Show()
			dur = time.Since(startTime)
		}
	}

	s.Fini()
}

func crash(s tcell.Screen, c Coordinates) {
	red := tcell.StyleDefault.Blink(true).Foreground(tcell.ColorRed)

	s.SetCell(c.x, c.y, red, 'X')
}
