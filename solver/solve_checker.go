package solver

import "github.com/Bryley/sudoku-solver/objects"

func IsSolved(board *objects.Board) bool {
	// Check Columns and Rows
	for index1 := uint8(0); index1 < 9; index1++ {
		rowNums := map[uint8]int{}
		colNums := map[uint8]int{}

		for index2 := uint8(0); index2 < 9; index2++ {
			rowNums[board.GetSquareNumber(index2, index1)]++
			colNums[board.GetSquareNumber(index1, index2)]++
		}
		if !mapContainsUniqueValues(rowNums) || !mapContainsUniqueValues(colNums) {
			return false
		}
	}

	// Check Quadrants
	for y := uint8(0); y < 3; y++ {
		for x := uint8(0); x < 3; x++ {
			quadValues := board.GetQuadrantNumbers(x, y)
			quadValuesMap := map[uint8]int{}

			for _, quadValue := range quadValues {
				quadValuesMap[quadValue]++
			}

			if !mapContainsUniqueValues(quadValuesMap) {
				return false
			}
		}
	}

	return true
}

func mapContainsUniqueValues(freq map[uint8]int) bool {
	for num, count := range freq {
		if num < 1 || num > 9 || count != 1 {
			return false
		}
	}
	return true
}
