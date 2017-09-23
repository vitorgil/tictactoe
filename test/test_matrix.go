package test

import "tictactoe/utils"
import "fmt"

func testRotateMatrix() {
	matrix := [3][3]rune {
		{' ', 'O', ' ',},
		{'X', ' ', ' ',},
		{'O', ' ', 'X',},
	}
	
	rotated := utils.Rotate(matrix, 90)

	controlMatrix := [3][3]rune {
		{'O', 'X', ' ',},
		{' ', ' ', 'O',},
		{'X', ' ', ' ',},
	}

	if rotated != controlMatrix {
		fmt.Printf("rotated: %c", rotated)
		panic("testRotateMatrix failed")
	}
}

func testRotateCell() {
	result_r, result_c := utils.CalculateRotatedCell(1, 2, 90)
	if result_r != 2 || result_c != 1 {
		panic("testRotateCell failed")
	}

	result_r, result_c = utils.CalculateRotatedCell(0, 2, 90)
	if result_r != 2 || result_c != 2 {
		panic("testRotateCell failed")
	}
}

func TestMatrix() {

	testRotateMatrix()

	testRotateCell()
}
