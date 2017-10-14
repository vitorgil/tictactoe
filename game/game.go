package game

import (
	"tictactoe/utils"
	"fmt"
	"github.com/eiannone/keyboard"
)

// Play starts the game!
func Play() {
	
	var g Game
	g.initialize()

	for !g.finished() {
		g.nextRound()
	}
}
	
/* Definition of Game
 */

// Game is your favourite game
type Game struct {
	panel _Panel
	p1    player
	p2    player

	// the player that is playing
	current *player
}

func (g *Game) initialize() {
	g.panel.initialize()

	// initialize players
	g.p1 = player{'X', "1"}
	g.p2 = player{'O', "2"}

	g.current = nil

	// playing via the console. Let's show something nice
	utils.ClearConsole()
	println()
	println("Welcome and Enjoy!")
	println("Press any key to start!")
	keyboard.GetSingleKey()
}

func (g Game) getPlayerWithSymbol(symbol rune) *player {
	if g.p1.symbol == symbol {
		return &g.p1
	} else if g.p2.symbol == symbol {
		return &g.p2
	}
	return nil
}

func (g Game) hasWinner() *player {

	matrix := &g.panel.matrix
	if !g.panel.isCellFree(cell{0, 0}) {
		if matrix[0][0] == matrix[0][1] && 
		matrix[0][0] == matrix[0][2] {
			return g.getPlayerWithSymbol(matrix[0][0])
		}
	}
	if !g.panel.isCellFree(cell{1, 0}) {
		if matrix[1][0] == matrix[1][1] && 
		matrix[1][0] == matrix[1][2] {
			return g.getPlayerWithSymbol(matrix[1][0])
		}
	}
	if !g.panel.isCellFree(cell{2, 0}) {
		if matrix[2][0] == matrix[2][1] && 
		matrix[2][0] == matrix[2][2] {
			return g.getPlayerWithSymbol(matrix[2][0])
		}
	}
	if !g.panel.isCellFree(cell{0, 0}) {
		if matrix[0][0] == matrix[1][0] && 
		matrix[0][0] == matrix[2][0] {
			return g.getPlayerWithSymbol(matrix[0][0])
		}
	}
	if !g.panel.isCellFree(cell{0, 1}) {
		if matrix[0][1] == matrix[1][1] && 
		matrix[0][1] == matrix[2][1] {
			return g.getPlayerWithSymbol(matrix[0][1])
		}
	}
	if !g.panel.isCellFree(cell{0, 2}) {
		if matrix[0][2] == matrix[1][2] && 
		matrix[0][2] == matrix[2][2] {
			return g.getPlayerWithSymbol(matrix[0][2])
		}
	}
	if !g.panel.isCellFree(cell{0, 0}) {
		if matrix[0][0] == matrix[1][1] && 
		matrix[0][0] == matrix[2][2] {
			return g.getPlayerWithSymbol(matrix[0][0])
		}
	}
	if !g.panel.isCellFree(cell{0, 2}) {
		if matrix[0][2] == matrix[1][1] && 
		matrix[0][2] == matrix[2][0] {
			return g.getPlayerWithSymbol(matrix[0][2])
		}
	}
	return nil
}

/* For now, the game is finished when all cells are non empty
 */
func (g *Game) finished() bool {
	// check if there is a winner
	pl := g.hasWinner()
	if pl != nil {
		println("Player", pl.name, "wins!")
		return true
	}

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if g.panel.isCellFree(cell{r, c}) {
				return false
			}
		}
	}
	return true
}

func (g *Game) assignNextPlayer() {
	if g.current == nil {
		g.current = &g.p1
	} else if g.current == &g.p1 {
		g.current = &g.p2
	} else {
		g.current = &g.p1
	}
}

func (g Game) showState() {
	
	utils.ClearConsole()

	// header sort of
	fmt.Println("\nPlayer", g.current.name, "paying!")
	fmt.Println()

	// the panel itself
	g.panel.print()
}
	
// performs the tasks necessary for the next round to happen
func (g *Game) nextRound() {
	g.assignNextPlayer()

	// Put this player's symbol on the first free cell
	g.panel.pointer = *g.panel.findFirstFreeCell(cell{0, 0})
	g.panel.assignCellValue(g.panel.pointer, g.current.symbol)

	// show the panel
	g.showState()

	// Wait until user hits some key
	c, _, _ := keyboard.GetSingleKey()

	for !utils.IsEnter(c) {
		// translate key to direction
		direction := getDirectionFromKey(c)

		// execute next move
		nextMove(&g.panel, direction, g.current)

		// show the game
		g.showState()

		// get next key
		c, _, _ = keyboard.GetSingleKey()
	}
}
	