package solver

import (
	"github.com/Bryley/sudoku-solver/objects"
)

func SolveBoard(board *objects.Board) {
    for {
        solvedOne := false
        for y := uint8(0); y < 9; y++ {
            for x := uint8(0); x < 9; x++ {
                num := board.GetSquareNumber(x, y)
                if num == 0 {
                    values := findPossibleValues(board, x, y)

                    if len(values) == 1 {
                        board.SetSquareNumber(x, y, values[0])
                        // fmt.Printf("Set square %v %v to %v\n", x, y, values[0])
                        solvedOne = true
                    }
                }
            }
        }
        if solvedOne == false {
            break
        }
    }
}


func findPossibleValues(board *objects.Board, x uint8, y uint8) []uint8 {
    seenValues := map[uint8]bool{}

    // Check index and column
    for index := uint8(0); index < 9; index++ {
        num := board.GetSquareNumber(x, index)
        seenValues[num] = true
        num = board.GetSquareNumber(index, y)
        seenValues[num] = true
    }

    // Check quadrant
    quadValues := board.GetQuadrantNumbers(x / 3, y / 3)
    for _, value := range quadValues {
        seenValues[value] = true
    }

    var possibleValues []uint8
    // Find possible values
    for i := uint8(1); i <= 9; i++ {
        _, ok := seenValues[i]
        if !ok {
            possibleValues = append(possibleValues, i)
        }
    }

    return possibleValues
}
