package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math/rand"
)

type Ball struct {
	speed  int
	x      int
	y      int
	deltaX int
	deltaY int
}

const ballSpeedMax = 5
const ballSpeedMin = 3

func randInRange(min, max int) int {
	number := rand.Intn(max-min) + min
	for number < ballSpeedMin && number > -ballSpeedMin {
		number = rand.Intn(max-min) + min
	}
	return number
}

var (
	ball = Ball{
		speed:  ballSpeedMax,
		x:      screenWidth / 2,
		y:      screenHeight / 2,
		deltaX: randInRange(-ballSpeedMax, ballSpeedMax),
		deltaY: randInRange(-ballSpeedMax, ballSpeedMax),
	}
)

func (b *Ball) reset() {
	b.x = screenWidth / 2
	b.y = screenHeight / 2
	b.deltaX = randInRange(-ballSpeedMax, ballSpeedMax)
	b.deltaY = randInRange(-ballSpeedMax, ballSpeedMax)
}

func (b *Ball) Update(g GameState) {
	b.x += b.deltaX
	b.y += b.deltaY
}

//go:embed assets/ball.png
var ballImageRaw []byte

func (b *Ball) Draw(screen *ebiten.Image) {
	ballImage, _, _ := ebitenutil.NewImageFromReader(bytes.NewReader(ballImageRaw))
	const offset = -20
	if ballImage != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(b.x+offset), float64(b.y+offset))
		screen.DrawImage(ballImage, op)
	}
}

func (b *Ball) ScoredLeft() bool {
	return ball.x <= 0
}

func (b *Ball) ScoredRight() bool {
	return ball.x+ballSize >= screenWidth
}
