package game

import (
	"container/ring"
	"fmt"
	"tictactoe/utils"
)

type _Panel struct {
	matrix     [3][3]rune // [row][col]
	pointer    cell
	horizontal *ring.Ring
	vertical   *ring.Ring
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

// seems to be necessary to pass pointer to pointer, otherwise
// the changes won't apply to the actual ring passed
func moveRingToPosition(pos int, r **ring.Ring) {
	for pos != (*r).Value.(int) {
		
		*r = (*r).Next()
	}
	utils.Debug("ring final position: %d\n", (*r).Value.(int))
}

/* Finds the first free cell after initialPosition
 */
func (p* _Panel) findFirstFreeCell(initialPosition cell) *cell {

	utils.DebugString("finding first free cell\n")
	utils.Debug("matrix: %c\n", p.matrix)
	utils.Debug("initial position: %v\n", initialPosition)

	// Move rings to the right place
	moveRingToPosition(initialPosition.col, &p.horizontal)
	moveRingToPosition(initialPosition.row, &p.vertical)
	
	// Find first free cell
	for vIndex := 0; vIndex < p.vertical.Len(); vIndex++ {
		row := p.vertical.Value.(int)
		for hIndex := 0; hIndex < p.horizontal.Len(); hIndex++ {
			col := p.horizontal.Value.(int)
			c := cell{row,col}
			utils.Debug("trying cell: %v\n", c)
	
			if p.isCellFree(c) {
				return &c
			}
			p.horizontal = p.horizontal.Next()
		}
		p.vertical = p.vertical.Next()
	}
	
	// return input. should happen only at the end
	return &initialPosition
}

/* Finds the first free cell after initialPosition
 */
func (p _Panel) findFirstFreeCellInDirection(initialPosition cell, d Direction) *cell {
	
	var rotation int
	if d == up {
		rotation = 90
	} else if d == down {
		rotation = -90
	} else if d == left {
		rotation = 180
	} else if d == right {
		rotation = 0
	}

	// get the rotated matrix from the panel's matrix and also the rotated initial position
	rotated := utils.Rotate(p.matrix, rotation)
	initialPositionRotated_r, initialPositionRotated_c := utils.CalculateRotatedCell(initialPosition.row, initialPosition.col, rotation)
	
	utils.Debug("rotated matrix : %c\n", rotated)
	utils.Debug("rotated initial position: %v\n", cell{initialPositionRotated_r, initialPositionRotated_c})

	// create new panel in which the matrix is the rotated one
	newPanel := p
	newPanel.matrix = rotated

	// for the new panel, find the first free cell from the initial position given
	cell_ := newPanel.findFirstFreeCell(cell{initialPositionRotated_r, initialPositionRotated_c})
	utils.Debug("first free cell rotated: %v\n", *cell_)

	// convert the found free position to the original matrix's coordinates
	// that is, rotate -90 degrees
	r, c := utils.CalculateRotatedCell(cell_.row, cell_.col, -rotation)
	
	utils.Debug("first free position final: %v\n", cell{r, c})
	utils.WaitInput()
	
	return &cell{r, c}
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
			fmt.Printf("%c", p.matrix[r][c])
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

	newCell := *p.findFirstFreeCellInDirection(p.pointer, d)
	
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
