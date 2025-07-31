package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"os"
	"time"
)

const (
	screenWidth  = 640
	screenHeight = 480
	playerWidth  = 10
	playerHeight = 110
	ballSize     = 15
	playerSpeed  = 5
	maxScore     = 10
)

type GameState string

const (
	MenuScreen         GameState = "menu"
	InstructionsScreen GameState = "instructions"
	Running            GameState = "running"
	Paused             GameState = "paused"
	Over               GameState = "over"
	ClosingScreen      GameState = "blank"
)

var gameState = MenuScreen

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
	if gameState == MenuScreen {
		menu.Update()
		currentDisplayedMessage = ""
		return nil
	}

	if gameState == InstructionsScreen {
		instructions.Update()
		currentDisplayedMessage = ""
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
			gameState = MenuScreen
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

	if gameState == ClosingScreen {
		currentDisplayedMessage = ""
		time.Sleep(200 * time.Millisecond)
		os.Exit(0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if gameState == ClosingScreen {
		screen.Fill(color.RGBA{0, 0, 0, 0xff})
		return
	}

	if gameState == MenuScreen {
		menu.Draw(screen)
		return
	}

	if gameState == InstructionsScreen {
		instructions.Draw(screen)
		return
	}

	screen.Fill(color.RGBA{0, 0, 0, 0xff})

	mplusNormalFont := LoadFont()
	DrawVerticallyHorizontallyCenteredText(screen, currentDisplayedMessage, mplusNormalFont, messageColourMain, defaultColorGlow)

	player1.Draw(screen)
	player2.Draw(screen)

	ball.Draw(screen)

	// Draw scores
	DrawLeftCenteredTopText(screen, fmt.Sprintf("%d", player1.score), mplusNormalFont, player1ColorMain, player1ColorGlow)
	DrawRightCenteredTopText(screen, fmt.Sprintf("%d", player2.score), mplusNormalFont, player2ColorMain, player2ColorGlow)
}
