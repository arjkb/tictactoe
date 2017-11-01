package tictactoe

import "strings"

func foo(num int) int {
	return num * num
}

func IsValidBoard(b string) bool {
	board := strings.ToUpper(b)

	if !hasValidCharsOnly(board) {
		return false
	}

	if len(board) != 11 {
		return false
	}

	return true
}

func hasValidCharsOnly(board string) bool {

	for i, ch := range board {
		if i == 3 || i == 7 {
			if ch != '|' {
				// fmt.Printf("A %d %c %s", i, ch, board)
				return false
			}
		} else {
			if ch != 'X' && ch != 'O' && ch != '-' {
				// fmt.Printf("B %d %c %s", i, ch, board)
				return false
			}
		}
	}

	return true
}
