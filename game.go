package main

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
