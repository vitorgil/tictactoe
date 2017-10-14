package main

import (
	"tictactoe/game"
	"tictactoe/test"
	"tictactoe/utils"
)

func main() {

	utils.ActivateDebug(false)

	// run tests, for now done before starting the game
	test.TestMatrix()

	game.Play()
}
