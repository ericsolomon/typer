package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type menuState struct {
	selectedItem int
	items        []*text.Text
}

func (s *menuState) enter(state) {
	if len(s.items) == 0 {
		basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

		for i, caption := range []string{
			"Start",
			"High Scores",
			"Quit",
		} {
			s.items = append(s.items, text.New(pixel.V(0, 0), basicAtlas))
			s.items[i].WriteString(caption)
		}
	}
}

func (*menuState) leave() {}

func (s *menuState) update(win *pixelgl.Window) state {
	var nextState state = menu

	if win.JustPressed(pixelgl.KeyEscape) {
		win.SetClosed(true)
	}

	if win.JustPressed(pixelgl.KeyDown) || win.JustPressed(pixelgl.KeyJ) {
		if s.selectedItem != len(s.items)-1 {
			s.selectedItem += 1
		}
	}

	if win.JustPressed(pixelgl.KeyUp) || win.JustPressed(pixelgl.KeyK) {
		if s.selectedItem > 0 {
			s.selectedItem -= 1
		}
	}

	if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.KeyKPEnter) {
		switch s.selectedItem {
		case 0:
			// nextState = play
			win.SetClosed(true)
		case 1:
			// nextState = scores
			win.SetClosed(true)
		case 2:
			win.SetClosed(true)
		}
	}

	//render menu items and selection
	const textScale = 4
	for i, item := range s.items {
		rectHeight := item.Bounds().H()
		m := pixel.IM.
			Moved(pixel.ZV.Sub(item.Bounds().Center())).
			Scaled(pixel.ZV, textScale).
			Moved(win.Bounds().Center()).
			Moved(pixel.V(0, -textScale*rectHeight*(float64(i)-float64(len(s.items))/2)))

		// draw rect around selected item
		if i == s.selectedItem {
			im := imdraw.New(nil)
			im.Color = pixel.RGB(0, 0.5, 0.5)
			r := item.Bounds()
			im.Push(
				m.Project(r.Min).Add(pixel.V(-20, 0)),
				m.Project(r.Max).Add(pixel.V(20, 0)),
			)
			im.Rectangle(0)
			im.Draw(win)
		}
		item.Draw(win, m)
	}

	return nextState
}
