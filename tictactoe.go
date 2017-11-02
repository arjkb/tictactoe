package tictactoe

import (
	"fmt"
	"strings"
)

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

func IsWinnable(board string, ch byte, indices []int) (bool, error) {
	var count int

	for _, index := range indices {
		if !isValidIndex(index) {
			return false, fmt.Errorf("IsWinnable: index %d is outside the board")
		}

		if board[index] == ch {
			count++
		}
	}

	if count == 2 {
		return true, nil
	} else {
		return false, nil
	}
}

func hasValidCharsOnly(board string) bool {
	for i, ch := range board {
		if i == 3 || i == 7 {
			if ch != '|' {
				return false
			}
		} else {
			if ch != 'X' && ch != 'O' && ch != '-' {
				return false
			}
		}
	}
	return true
}

func isValidIndex(i int) bool {
	// 0123456789X
	// ---|---|---
	if i >= 0 && i <= 10 {
		return true
	} else {
		return false
	}
}

func isIndicesDifferent(idx [3]int) bool {
	if idx[0] == idx[1] || idx[0] == idx[2] || idx[1] == idx[2] {
		return false
	} else {
		return true
	}
}
