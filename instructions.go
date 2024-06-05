package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

type Instructions struct {
	regularFont font.Face
}

func (m *Instructions) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		gameState = MenuScreen
	}
}

func (m *Instructions) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	text.Draw(screen, "Instructions", m.regularFont, 50, 30, color.White)
	text.Draw(screen, "Player 1: W and S", m.regularFont, 50, 70, color.White)
	text.Draw(screen, "Player 2: Up and Down", m.regularFont, 50, 110, color.White)
	text.Draw(screen, "Pause: P", m.regularFont, 50, 150, color.White)
	text.Draw(screen, "Back to menu from here: Esc", m.regularFont, 50, 200, color.White)
}

var instructions = &Instructions{
	regularFont: LoadFont(),
}
