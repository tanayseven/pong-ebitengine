package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
	"os"
)

type MenuState string

type Menu struct {
	selection   MenuState
	regularFont font.Face
}

const (
	GameStart        MenuState = "Start Game"
	GameInstructions MenuState = "Instructions"
	GameExit         MenuState = "Exit Game"
)

func (m *Menu) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		menu.NextSelection()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		menu.PrevSelection()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if m.selection == GameStart {
			gameState = Running
		} else if m.selection == GameInstructions {
			gameState = InstructionsScreen
		} else if m.selection == GameExit {
			os.Exit(0)
		}
	}
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	// draw menu items
	screen.Fill(color.RGBA{0, 0, 0, 255})
	selectColor := color.RGBA{255, 0, 0, 255}

	startColor := color.RGBA{255, 255, 255, 255}
	instructionsColor := color.RGBA{255, 255, 255, 255}
	quitColor := color.RGBA{255, 255, 255, 255}

	switch m.selection {
	case GameStart:
		startColor = selectColor
		break
	case GameInstructions:
		instructionsColor = selectColor
		break
	case GameExit:
		quitColor = selectColor
		break
	}

	text.Draw(screen, "Start Game", m.regularFont, 50, 30, startColor)
	text.Draw(screen, "Instructions", m.regularFont, 50, 70, instructionsColor)
	text.Draw(screen, "Quit", m.regularFont, 50, 110, quitColor)
}

func (m *Menu) NextSelection() {
	switch m.selection {
	case GameStart:
		m.selection = GameInstructions
		break
	case GameInstructions:
		m.selection = GameExit
		break
	case GameExit:
		m.selection = GameStart
		break
	}
}

func (m *Menu) PrevSelection() {
	switch m.selection {
	case GameStart:
		m.selection = GameExit
		break
	case GameInstructions:
		m.selection = GameStart
		break
	case GameExit:
		m.selection = GameInstructions
		break
	}
}

var (
	menu = &Menu{
		selection:   GameStart,
		regularFont: LoadFont(),
	}
)
