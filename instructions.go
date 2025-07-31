package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"image/color"
)

type Instructions struct {
	retroFont font.Face
}

func (i *Instructions) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		gameState = MenuScreen
	}
}

func (i *Instructions) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	defaultColor := menuColorMain
	spacing := 40

	texts := []struct {
		Text      string
		Color     color.Color
		GlowColor color.Color
	}{
		{"Instructions", defaultColor, defaultColorGlow},
		{"Player 1: [W] and [S]", defaultColor, defaultColorGlow},
		{"Player 2: [Up] and [Down]", defaultColor, defaultColorGlow},
		{"Pause: [P]", defaultColor, defaultColorGlow},
		{"Back to menu from here: [Esc]", defaultColor, defaultColorGlow},
	}

	initialY := screenHeight/2 - len(texts)*spacing/2

	for n, t := range texts {
		y := initialY + n*spacing
		DrawHorizontallyCenteredText(screen, t.Text, y, i.retroFont, t.Color, t.GlowColor)
	}
}

var instructions = &Instructions{
	retroFont: LoadFont(),
}
