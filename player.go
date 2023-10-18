package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Player struct {
	x       int
	y       int
	score   int
	upKey   ebiten.Key
	downKey ebiten.Key
}

var (
	player1 = &Player{
		x:       0,
		y:       screenHeight / 2,
		score:   0,
		upKey:   ebiten.KeyW,
		downKey: ebiten.KeyS,
	}
	player2 = &Player{
		x:       screenWidth - playerWidth,
		y:       screenWidth / 2,
		score:   0,
		downKey: ebiten.KeyDown,
		upKey:   ebiten.KeyUp,
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
	vector.DrawFilledRect(screen, float32(p.x), float32(p.y), float32(playerWidth), float32(playerHeight), color.White, true)
}

func (p *Player) ScoredGoal() {
	p.score++
}
