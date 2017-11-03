package tictactoe

import (
	"fmt"
	"strings"
)

func GetEmptyBoard() string {
	var emptyBoard string = "---|---|---"
	return emptyBoard
}

func IsValidBoard(b string) bool {
	board := strings.ToUpper(b)
	const EXPECTEDLENGTH = 9 + 2

	if !hasValidCharsOnly(board) || len(board) != EXPECTEDLENGTH {
		return false
	}

	return true
}

func IsWinnable(board string, my_symbol byte, indices [3]int) (bool, []int, error) {
	var count_mysym, count_oppsym int

	opp_symbol := getOpponentSymbol(my_symbol)

	for _, index := range indices {
		if !isValidIndex(index) {
			return false, nil, fmt.Errorf("IsWinnable: index %d is outside the board")
		}

		switch board[index] {
		case my_symbol:
			count_mysym++
		case opp_symbol:
			count_oppsym++
		}
	}

	if count_mysym == 2 && count_oppsym == 0 {
		return true, indices[:], nil
	}
	return false, nil, nil
}

func MakeWinMove(board string, move [3]int, symbol byte) (string, error) {
	boardBytes := []byte(board)

	winnable, _, _ := IsWinnable(board, symbol, move)
	if !winnable {
		return "", fmt.Errorf("MakeWinMove() IsWinnable(%v, %q, %v)=%v", board, symbol, move, winnable)
	}

	pos, err := getEmptyPos(board, move[:])
	if err != nil {
		return "", fmt.Errorf("MakeWinMove(): %v", err)
	}

	boardBytes[pos] = symbol

	return string(boardBytes), nil
}

func BlockWinMove(board string, move [3]int, symbol byte) (string, error) {
	boardBytes := []byte(board)

	// check if the caller's opponent can win
	winnable, _, _ := IsWinnable(board, getOpponentSymbol(symbol), move)
	if !winnable {
		return "", fmt.Errorf("BlockWinMove() IsWinnable(%v, %q, %v)=%v", board, symbol, move, winnable)
	}

	pos, err := getEmptyPos(board, move[:])
	if err != nil {
		return "", fmt.Errorf("BlockWinMove(): %v", err)
	}

	boardBytes[pos] = symbol

	return string(boardBytes), nil
}

func MakeRandomMove(board string, move []int, symbol byte) (string, error) {
	boardBytes := []byte(board)

	pos, err := getEmptyPos(board, move)
	if err != nil {
		return "", fmt.Errorf("MakeRandomMove(): %v", err)
	}

	boardBytes[pos] = symbol

	return string(boardBytes), nil
}

func getEmptyPos(board string, indices []int) (int, error) {
	for _, index := range indices {
		if board[index] == '-' {
			return index, nil
		}
	}

	return 0, fmt.Errorf("getEmptyPos() board:%v. No positions in %v are empty", board, indices)
}

func HasWon(b string, symbol byte) bool {
	//check if somebody has won

	patterns := [][]int{
		// horizontals
		{0, 1, 2},
		{4, 5, 6},
		{8, 9, 10},

		// verticals
		{0, 4, 8},
		{1, 5, 9},
		{2, 6, 10},

		// diagonals
		{0, 5, 10},
		{2, 5, 8},
	}

	for _, index := range patterns {
		if b[index[0]] == symbol && b[index[1]] == symbol && b[index[2]] == symbol {
			// won!
			return true
		}
	}

	return false
}

func getOpponentSymbol(ch byte) byte {
	switch ch {
	case 'X':
		return 'O'
	case 'O':
		return 'X'
	default:
		return '-'
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
