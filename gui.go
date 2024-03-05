package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

func DrawCenteredText(screen *ebiten.Image, str string, cx, cy int, font font.Face, clr color.Color) {
	bounds := text.BoundString(font, str)
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2
	text.Draw(screen, str, font, x, y, clr)
}

var currentDisplayedMessage = ""

const gameStartMessage = "Press space to start the game"
const gamePausedMessage = "Game paused, press P to continue"
const player1WonMessage = "Player 1 won, press space to restart"
const player2WonMessage = "Player 2 won, press space to restart"
