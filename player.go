package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	x           int
	y           int
	score       int
	upKey       ebiten.Key
	downKey     ebiten.Key
	paddleImage []byte
}

//go:embed assets/paddle1.png
var paddle1ImageRaw []byte

//go:embed assets/paddle2.png
var paddle2ImageRaw []byte

var (
	player1 = &Player{
		x:           10,
		y:           screenHeight / 2,
		score:       0,
		upKey:       ebiten.KeyW,
		downKey:     ebiten.KeyS,
		paddleImage: paddle1ImageRaw,
	}
	player2 = &Player{
		x:           screenWidth - playerWidth - 10,
		y:           screenWidth / 2,
		score:       0,
		downKey:     ebiten.KeyDown,
		upKey:       ebiten.KeyUp,
		paddleImage: paddle2ImageRaw,
	}
)

func (p *Player) Update(g GameState) {
	if g == Paused {
		return
	}
	if ebiten.IsKeyPressed(p.upKey) && p.y > 0 {
		p.y -= playerSpeed
	}
	if ebiten.IsKeyPressed(p.downKey) && p.y+playerHeight < screenHeight {
		p.y += playerSpeed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	const offset = -20
	paddleImage, _, _ := ebitenutil.NewImageFromReader(bytes.NewReader(p.paddleImage))
	if paddleImage != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(p.x+offset), float64(p.y+offset))
		screen.DrawImage(paddleImage, op)
	}
}

func (p *Player) ScoredGoal() {
	p.score++
}
