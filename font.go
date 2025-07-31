package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

//go:embed assets/neon-spark.otf
var ticketingFont []byte

func LoadFont() font.Face {
	tt, _ := opentype.Parse(ticketingFont)
	mplusNormalFont, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return mplusNormalFont
}

var offsets = []struct{ dx, dy int }{
	{-2, 0}, {2, 0}, {0, -2}, {0, 2},
	{-2, -2}, {-2, 2}, {2, -2}, {2, 2},
}

func DrawHorizontallyCenteredText(screen *ebiten.Image, msg string, y int, face font.Face, textColor color.Color, glowColor color.Color) {
	// Calculate the width of the text
	bounds := text.BoundString(face, msg)
	textWidth := bounds.Dx()

	// Center the text
	screenWidth, _ := screen.Size()
	x := (screenWidth - textWidth) / 2

	// Draw multiple layers of the same text with offset and semi-transparent color to simulate glow
	for _, o := range offsets {
		text.Draw(screen, msg, face, x+o.dx, y+o.dy, glowColor)
	}

	// Draw actual text on top
	text.Draw(screen, msg, face, x, y, textColor)
}

func DrawVerticallyHorizontallyCenteredText(screen *ebiten.Image, msg string, face font.Face, textColor color.Color, glowColor color.Color) {

	_, screenHeight := screen.Size()
	y := screenHeight / 2

	DrawHorizontallyCenteredText(screen, msg, y, face, textColor, glowColor)
}

func DrawLeftCenteredTopText(screen *ebiten.Image, msg string, face font.Face, textColor color.Color, shadowColor color.Color) {
	// Calculate the width of the text
	bounds := text.BoundString(face, msg)
	textWidth := bounds.Dx()

	// Center the text
	screenWidth, _ := screen.Size()
	x := ((screenWidth / 2) - textWidth) / 2
	y := 30

	// Draw multiple layers of the same text with offset and semi-transparent color to simulate glow
	for _, o := range offsets {
		text.Draw(screen, msg, face, x+o.dx, y+o.dy, shadowColor)
	}

	// Draw actual text on top
	text.Draw(screen, msg, face, x, y, textColor)
}

func DrawRightCenteredTopText(screen *ebiten.Image, msg string, face font.Face, textColor color.Color, shadowColor color.Color) {
	// Calculate the width of the text
	bounds := text.BoundString(face, msg)
	textWidth := bounds.Dx()

	// Center the text vertically
	screenWidth, _ := screen.Size()
	x := ((screenWidth/2)-textWidth)/2 + (screenWidth / 2)
	y := 30

	// Draw multiple layers of the same text with offset and semi-transparent color to simulate glow
	offsets := []struct{ dx, dy int }{
		{-2, 0}, {2, 0}, {0, -2}, {0, 2},
		{-2, -2}, {-2, 2}, {2, -2}, {2, 2},
	}
	for _, o := range offsets {
		text.Draw(screen, msg, face, x+o.dx, y+o.dy, shadowColor)
	}

	// Draw actual text on top
	text.Draw(screen, msg, face, x, y, textColor)
}
