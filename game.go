package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

const (
	screenWidth  = 640
	screenHeight = 480
	playerWidth  = 15
	playerHeight = 100
	ballSize     = 15
	playerSpeed  = 5
	maxScore     = 10
)

type GameState string

const (
	New     GameState = "new"
	Running GameState = "running"
	Paused  GameState = "paused"
	Over    GameState = "over"
)

var gameState = New

type Game struct {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

type GameObjects interface {
	Update(g GameState)
	Draw(screen *ebiten.Image)
}

func (g *Game) Update() error {
	if gameState == New {
		currentDisplayedMessage = gameStartMessage
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			gameState = Running
			currentDisplayedMessage = ""
		}
		return nil
	}

	if gameState == Paused {
		currentDisplayedMessage = gamePausedMessage
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			gameState = Running
			currentDisplayedMessage = ""
		}
		return nil
	}

	if gameState == Over {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			gameState = New
			player1.score = 0
			player2.score = 0
			currentDisplayedMessage = gameStartMessage
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		gameState = Paused
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyAlt) && inpututil.IsKeyJustPressed(ebiten.KeyF4) {
		return ebiten.Termination
	}

	player1.Update(gameState)
	player2.Update(gameState)

	// Ball movement
	ball.Update(gameState)

	// Ball collision with walls
	if ball.y <= 0 || ball.y+ballSize >= screenHeight {
		ball.deltaY = -ball.deltaY
	}

	// Ball collision with players
	if ball.x <= playerWidth && ball.y+ballSize >= player1.y && ball.y <= player1.y+playerHeight {
		ball.deltaX = -ball.deltaX
	}

	if ball.x+ballSize >= screenWidth-playerWidth && ball.y+ballSize >= player2.y && ball.y <= player2.y+playerHeight {
		ball.deltaX = -ball.deltaX
	}

	if ball.ScoredLeft() {
		player2.ScoredGoal()
		ball.reset()
	}

	if ball.ScoredRight() {
		player1.ScoredGoal()
		ball.reset()
	}

	if player1.score == maxScore {
		gameState = Over
		currentDisplayedMessage = player1WonMessage
	}

	if player2.score == maxScore {
		gameState = Over
		currentDisplayedMessage = player2WonMessage
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0xff})

	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	mplusNormalFont, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	DrawCenteredText(screen, currentDisplayedMessage, screenWidth/2, screenHeight/2, mplusNormalFont, color.White)

	player1.Draw(screen)
	player2.Draw(screen)

	// Draw ball
	ball.Draw(screen)

	// Draw scores
	DrawCenteredText(screen, fmt.Sprintf("%d", player1.score), screenWidth/4, 60, mplusNormalFont, color.White)
	DrawCenteredText(screen, fmt.Sprintf("%d", player2.score), screenWidth*3/4, 60, mplusNormalFont, color.White)
}
