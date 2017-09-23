package main

import "tictactoe/src"
import "tictactoe/test"
import "tictactoe/utils"

func main() {

	utils.ActivateDebug(false)

	// run tests, for now done before starting the game
	test.TestMatrix()

	tictactoe.Play()
}
