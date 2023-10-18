package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

const (
	screenWidth  = 640
	screenHeight = 480
	playerWidth  = 15
	playerHeight = 100
	ballSize     = 15
	playerSpeed  = 5
)

type GameState string

const (
	Running GameState = "running"
	Paused  GameState = "paused"
)

var gameState = Running

// Game stuff
var gameOverMessage = ""

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

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0xff})

	player1.Draw(screen)
	player2.Draw(screen)

	// Draw ball
	ball.Draw(screen)

	// Draw scores
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player 1: %d\nPlayer 2: %d\n%s", player1.score, player2.score, gameOverMessage))
}
