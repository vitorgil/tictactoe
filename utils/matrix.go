package utils


// Rotate rotates matrix by angle degrees
func Rotate(matrix [3][3]rune, angle int) [3][3]rune {
	
	rotated := matrix
	// if the rotation is other than the cases below, like 0 degrees,
	// then the input is returned.
	
	if angle == 90 {
		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				new_r, new_c := CalculateRotatedCell(r, c, 90)
				rotated[new_r][new_c] = matrix[r][c]
			}
		}
	} else if angle == -90 {
		rotated = Rotate(matrix, 180)
		rotated = Rotate(rotated, 90)
	} else if angle == 180 || angle == -180 {
		rotated = Rotate(matrix, 90)
		rotated = Rotate(rotated, 90)
	}

	return rotated
}

// assumes that a matrix
// 	- is square
//  - is 3 dimensional
func CalculateRotatedCell(r, c, angle int) (int, int) {

	DebugString("Calculating rotated cell\n")
	Debug("position: %d %d\n", r, c)
	Debug("angle: %d\n", angle)
	
	rotated_r, rotated_c := r, c
	// if the rotation is other than the cases below, like 0 degrees,
	// then the input is returned.
	if angle == 90 {
		rotated_r = c
		rotated_c = 2 - r
	} else if angle == -90 {
		rotated_r, rotated_c = CalculateRotatedCell(r, c, 180)
		rotated_r, rotated_c = CalculateRotatedCell(rotated_r, rotated_c, 90)
	} else if angle == 180 || angle == -180 {
		rotated_r, rotated_c = CalculateRotatedCell(r, c, 90)
		rotated_r, rotated_c = CalculateRotatedCell(rotated_r, rotated_c, 90)
	}

	Debug("rotated: %d %d\n", rotated_r, rotated_c)

	return rotated_r, rotated_c
}
