package main

import "github.com/faiface/pixel/pixelgl"

type gameState struct {
}

func (s *gameState) enter(from state) {

}

func (s *gameState) leave() {}

func (s *gameState) update(win *pixelgl.Window) state {
	var nextState state = game

	if win.JustPressed(pixelgl.KeyEscape) {
		nextState = menu
	}

	return nextState
}
