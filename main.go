package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong Game")
	updatableGameObjects := []GameObjects{}
	updatableGameObjects = append(updatableGameObjects, player1)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
