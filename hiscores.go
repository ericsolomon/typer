package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type scoreState struct {
	selectedItem int
	items        []*text.Text
}

func (s *scoreState) enter(from state) {
	if len(s.items) == 0 {
		basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

		for i, caption := range []string{
			"Back",
		} {
			s.items = append(s.items, text.New(pixel.V(0, 0), basicAtlas))
			s.items[i].WriteString(caption)
		}
	}
}

func (s *scoreState) leave() {}

func (s *scoreState) update(win *pixelgl.Window) state {
	var nextState state = score

	if win.JustPressed(pixelgl.KeyEscape) {
		nextState = menu
	}

	if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.KeyKPEnter) {
		switch s.selectedItem {
		case 0:
			nextState = menu
		}
	}

	//render menu items and selection
	const textScale = 4
	for i, item := range s.items {
		m := pixel.IM.
			Moved(pixel.ZV.Sub(item.Bounds().Center())).
			Scaled(pixel.ZV, textScale).
			Moved(pixel.V(win.Bounds().Center().X, win.Bounds().Min.X+40))

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
