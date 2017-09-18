package tictactoe

import (
	"fmt"
	"os"
	"os/exec"
	"tictactoe/utils"

	"github.com/eiannone/keyboard"
)

type _Panel struct {
	matrix  [3][3]rune // [row][col]
	pointer cell
}

func (p *_Panel) initialize() {
	for r := 0; r < 3; r = r + 1 {
		for c := 0; c < 3; c = c + 1 {
			p.matrix[r][c] = ' '
		}
	}
	p.pointer = cell{0, 0}
}

/* Finds the first free cell after initialPosition
 */
func (p _Panel) findFirstFreeCell(initialPosition cell) *cell {
	for r := initialPosition.row; r < 3; r = r + 1 {
		for col := initialPosition.col; col < 3; col = col + 1 {
			if p.matrix[r][col] == ' ' {
				return &cell{r, col}
			}
		}
	}
	panic("no free cells!")
}

func (p *_Panel) assignCellValue(c cell, val rune) {
	p.matrix[c.row][c.col] = val
}

/* Print a panel
 */
func (p _Panel) print() {
	fmt.Println()
	for r := 0; r < 3; r = r + 1 {
		fmt.Print(" ")
		for c := 0; c < 3; c = c + 1 {
			fmt.Printf("%q", p.matrix[r][c])
			if c < 3-1 {
				fmt.Print(" | ")
			}
		}
		fmt.Print("\n")
		if r < 3-1 {
			fmt.Println("-----------")
		}
	}
	fmt.Println()
}

/* player
 */
type player struct {
	symbol rune
	name   string
}

/* cell
 */
type cell struct {
	row int
	col int
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
	clearConsole()
	println()
	println("Welcome and Enjoy!")
	println("Press any key to start!")
	keyboard.GetSingleKey()
}

/* For now, the game is finished when all cells are non empty
 */
func (g *Game) finished() bool {
	for r := 0; r < 3; r = r + 1 {
		for c := 0; c < 3; c = c + 1 {
			if g.panel.matrix[r][c] == ' ' {
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

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (g Game) showState() {

	clearConsole()

	// header sort of
	fmt.Println("\nPlayer", g.current.name, "paying!")
	fmt.Println()

	// the panel itself
	g.panel.print()
}

func (g *Game) nextRound() {

	g.assignNextPlayer()

	// Put this player's symbol on the first free cell and show the panel
	g.panel.pointer = *g.panel.findFirstFreeCell(cell{0, 0})
	g.panel.assignCellValue(g.panel.pointer, g.current.symbol)

	g.showState()

	// Wait until user hits keys
	c, _, _ := keyboard.GetSingleKey()

	for !utils.IsEnter(c) {
		// translate key to direction
		direction := getDirectionFromKey(c)

		// execute next move
		nextMove(&g.panel, direction, g.current)

		g.showState()

		// get next key
		c, _, _ = keyboard.GetSingleKey()
	}

}

// Direction defines the types of movesthe user can make: Up, Down, etc
type Direction int8

const (
	undefined Direction = iota
	up
	down
	left
	right
)

func nextMove(p *_Panel, d Direction, pl *player) {
	oldCell := p.pointer
	switch d {
	case up:
		p.pointer.row--
	case down:
		p.pointer.row++
	case left:
		p.pointer.col--
	case right:
		p.pointer.col++
	default:
		return
	}
	p.pointer = *p.findFirstFreeCell(p.pointer)
	// assign new one
	p.assignCellValue(p.pointer, pl.symbol)
	// remove old one
	p.assignCellValue(oldCell, ' ')
}

func getDirectionFromKey(c rune) Direction {
	switch c {
	case 'w':
		return up
	case 's':
		return down
	case 'a':
		return left
	case 'd':
		return right
	default:
		return undefined
	}
}

// Play starts the game!
func Play() {

	var g Game
	g.initialize()

	for !g.finished() {
		g.nextRound()
	}
}
