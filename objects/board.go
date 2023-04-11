package objects

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Board struct {
	// 5 bytes per row
	squares [45]byte
}

func (b Board) GetSquareNumber(x uint8, y uint8) uint8 {
	squareByte := b.squares[(y*5)+x/2]

	if x%2 == 0 {
		return squareByte >> 4
	} else {
		return 0b1111 & squareByte
	}
}

func (b *Board) SetSquareNumber(x uint8, y uint8, value uint8) {
	index := (y * 5) + x/2
	squareByte := b.squares[index]

	var newByte byte

	if x%2 == 0 {
		newByte = (squareByte & 0b1111) | value<<4
	} else {
		newByte = (squareByte & (0b1111 << 4)) | value
	}
	b.squares[index] = newByte
}

func (b *Board) GetQuadrantNumbers(x uint8, y uint8) [9]uint8 {
	var values [9]uint8
	valuesIndex := 0

	for yOffset := 0; yOffset < 3; yOffset++ {
		for xOffset := 0; xOffset < 3; xOffset++ {
			values[valuesIndex] = b.GetSquareNumber(x*3+uint8(xOffset), y*3+uint8(yOffset))
			valuesIndex++
		}
	}
    return values
}

func NewBoardFromFile(fileName string) (*Board, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var numbers [81]uint8
	index := 0

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if unicode.IsDigit(char) {
			numbers[index] = uint8(char - '0')
			index++
			continue
		}
		if char == '#' {
			numbers[index] = 0
			index++
		}
	}
	board, err := NewBoard(numbers)
	return board, err
}

func (b Board) PrintBoard() {
	for y := uint8(0); y < 9; y++ {
		for x := uint8(0); x < 9; x++ {
			fmt.Print(b.GetSquareNumber(x, y))
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func NewBoard(numbers [81]uint8) (*Board, error) {
	var squares [45]byte

	for row := 0; row < 9; row++ {
		for i := 0; i < 9; i += 2 {
			index1 := i + row*9
			index2 := index1 + 1
			num1 := numbers[index1]
			num2 := func() uint8 {
				if i == 8 {
					return 0
				}
				return numbers[index2]
			}()
			var squareByte byte = num1<<4 | num2
			squares[row*5+(i/2)] = squareByte
		}
	}
	return &Board{
		squares: squares,
	}, nil
}
