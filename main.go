package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type state interface {
	enter(from state)
	update(win *pixelgl.Window) state
	leave()
}

// game states
var (
	menu  = &menuState{}
	game  = &gameState{}
	score = &scoreState{}
)

func run() {
	var state state = menu
	state.enter(nil)

	cfg := pixelgl.WindowConfig{
		Title:  "Typer",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if (err) != nil {
		panic(err)
	}
	win.SetCursorVisible(false)

	// basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	// basicText := text.New(pixel.V(0, 500), basicAtlas)

	// vel := 30.0
	// last := time.Now()

	for !win.Closed() {
		win.Clear(colornames.Black)
		// if menu.isActive {
		// 	menu.Draw(win)
		// 	menu.Update(win)
		// 	win.Update()
		// 	continue
		// }
		// dt := time.Since(last).Seconds()
		// last = time.Now()

		newState := state.update(win)
		if state != newState {
			state.leave()
			newState.enter(state)
		}
		state = newState
		// basicText.Clear()
		// basicText.Orig = pixel.V(basicText.Orig.X+(vel*dt), 500)
		// fmt.Fprintln(basicText, "foobar")
		// basicText.Draw(win, pixel.IM)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
