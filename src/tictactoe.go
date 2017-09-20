package tictactoe

import (
	"container/ring"
	"fmt"
	"os"
	"os/exec"
	"tictactoe/utils"

	"github.com/eiannone/keyboard"
)

type _Panel struct {
	matrix     [3][3]rune // [row][col]
	pointer    cell
	horizontal *ring.Ring
	vertical   *ring.Ring
}

func (p _Panel) getTransposed() [3][3]rune {
	var matrix [3][3]rune
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			matrix[r][c] = p.matrix[3-1-r][3-1-c]
		}
	}
	return matrix
}

func (p *_Panel) initialize() {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			p.emptyCell(cell{r, c})
		}
	}

	p.pointer = cell{0, 0}

	p.horizontal = ring.New(3)
	p.horizontal.Value = 0
	p.horizontal = p.horizontal.Next()
	p.horizontal.Value = 1
	p.horizontal = p.horizontal.Next()
	p.horizontal.Value = 2
	p.horizontal = p.horizontal.Next()

	p.vertical = ring.New(3)
	p.vertical.Value = 0
	p.vertical = p.vertical.Next()
	p.vertical.Value = 1
	p.vertical = p.vertical.Next()
	p.vertical.Value = 2
	p.vertical = p.vertical.Next()

}

/* Finds the first free cell after initialPosition
 */
func (p *_Panel) findFirstFreeCell(initialPosition cell) *cell {

	// Move rings to the right place
	for initialPosition.col != p.horizontal.Value {
		p.horizontal = p.horizontal.Next()
	}
	for initialPosition.row != p.vertical.Value {
		p.vertical = p.vertical.Next()
	}

	// Find first free cell
	for hIndex := 0; hIndex < p.horizontal.Len(); hIndex++ {
		hVal := p.horizontal.Value.(int)
		for vIndex := 0; vIndex < p.vertical.Len(); vIndex++ {
			vVal := p.vertical.Value.(int)
			if p.isCellFree(cell{vVal, hVal}) {
				// optimize making the cell
				return &cell{vVal, hVal}
			}
			p.vertical = p.vertical.Next()
		}
		p.horizontal = p.horizontal.Next()
	}

	panic("no free cells!")
}

func nextCell(c cell, d Direction) *cell {

	r := ring.New(3)
	r.Value = 0
	r = r.Next()
	r.Value = 1
	r = r.Next()
	r.Value = 2

	switch d {
	case up:
		for c.row != r.Value {
			r = r.Prev()
		}
		r = r.Prev()
		val, _ := r.Value.(int)
		return &cell{val, c.col}
	case down:
		for c.row != r.Value {
			r = r.Next()
		}
		r = r.Next()
		return &cell{r.Value.(int), c.col}
	case left:
		for c.col != r.Value {
			r = r.Prev()
		}
		r = r.Prev()
		return &cell{c.row, r.Value.(int)}
	case right:
		for c.col != r.Value {
			r = r.Next()
		}
		r = r.Next()
		return &cell{c.row, r.Value.(int)}
	}
	panic("problem in choosing next cell")
}

/* Finds the first free cell after initialPosition
 */
func (p _Panel) findFirstFreeCellInDirection(initialPosition cell, d Direction) *cell {
	if d == up {
		
		transposed := p.getTransposed()
		
		var newPanel _Panel
		newPanel.matrix = transposed
		newPanel.findFirstFreeCell(initialPosition)

		
	} else if d == down {
		for col := initialPosition.col; col < 3; col++ {
			var _cell cell
			for r := 0; r < 2; r++ {
				_cell = *nextCell(cell{initialPosition.row, col}, d)
				if p.isCellFree(_cell) {
					return &_cell
				}
			}
			// at the next round, we start from the top
			initialPosition.row = 0
		}
	} else if d == left {
		for r := initialPosition.row; r >= 0; r-- {
			var _cell cell
			for c := 0; c < 2; c++ {
				_cell = *nextCell(cell{r, initialPosition.col}, d)
				if p.isCellFree(_cell) {
					return &_cell
				}
			}
			// at the next round, we start from the right
			initialPosition.col = 2
		}
	} else if d == right {
		for r := initialPosition.row; r < 3; r++ {
			var _cell cell
			for c := 0; c < 2; c++ {
				_cell = *nextCell(cell{r, initialPosition.col}, d)
				if p.isCellFree(_cell) {
					return &_cell
				}
			}
			// at the next round, we start from the left
			initialPosition.col = 0
		}
	}
	panic("no free cells!")
}

func (p *_Panel) assignCellValue(c cell, val rune) {
	if c.row < 0 || c.row >= 3 || c.col < 0 || c.col >= 3 {
		return
	}
	p.matrix[c.row][c.col] = val
}

func (p *_Panel) emptyCell(c cell) {
	if c.row < 0 || c.row >= 3 || c.col < 0 || c.col >= 3 {
		return
	}
	p.matrix[c.row][c.col] = ' '
}

/* Print a panel
 */
func (p _Panel) print() {
	fmt.Println()
	for r := 0; r < 3; r++ {
		fmt.Print(" ")
		for c := 0; c < 3; c++ {
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

// Checks whether the cell given in p is free
func (p _Panel) isCellFree(c cell) bool {
	if c.row < 0 || c.row >= 3 || c.col < 0 || c.col >= 3 {
		return false
	}
	return p.matrix[c.row][c.col] == ' '
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

	// save the current cell
	//newCell := p.pointer

	newCell := *p.findFirstFreeCellInDirection(p.pointer, d)
	// update pointer

	if !p.isCellFree(newCell) {
		return
	}

	// assign new one
	p.assignCellValue(newCell, pl.symbol)

	// empty the old cell
	p.emptyCell(p.pointer)

	// find the first free cell from where pointer is
	p.pointer = newCell
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
}
