package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type Ball struct {
	speed  int
	x      int
	y      int
	deltaX int
	deltaY int
}

const ballSpeed = 4

var (
	ball = Ball{
		speed:  4,
		x:      screenWidth / 2,
		y:      screenHeight / 2,
		deltaX: ballSpeed,
		deltaY: ballSpeed,
	}
)

func (b *Ball) reset() {
	b.x = screenWidth / 2
	b.y = screenHeight / 2
	b.deltaX = ballSpeed
	b.deltaY = ballSpeed
}

func (b *Ball) Update(g GameState) {
	b.x += b.deltaX
	b.y += b.deltaY
}

func (b *Ball) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(b.x), float64(b.y), float64(ballSize), float64(ballSize), color.White)
}

func (b *Ball) ScoredLeft() bool {
	return ball.x <= 0
}

func (b *Ball) ScoredRight() bool {
	return ball.x+ballSize >= screenWidth
}
