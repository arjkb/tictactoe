package tictactoe

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	CLIENTWON    = "client won"
	SERVERWON    = "server won"
	CLIENTSYMBOL = 'X'
	SERVERSYMBOL = 'O'
	TIE          = "tie"
)

var WinPatterns [][]int
var AllSquares []int

func init() {
	WinPatterns = [][]int{
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

	AllSquares = []int{0, 1, 2, 4, 5, 6, 8, 9, 10}
}

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

func IsWinnable(board string, my_symbol byte, indices [3]int) (bool, error) {
	var count_mysym, count_oppsym int

	opp_symbol := getOpponentSymbol(my_symbol)

	for _, index := range indices {
		if !isValidIndex(index) {
			return false, fmt.Errorf("IsWinnable: index %d is outside the board")
		}

		switch board[index] {
		case my_symbol:
			count_mysym++
		case opp_symbol:
			count_oppsym++
		}
	}

	if count_mysym == 2 && count_oppsym == 0 {
		return true, nil
	}
	return false, nil
}

func CanWinNext(board string, symbol byte) (bool, [3]int) {
	var parr [3]int
	for _, pslice := range WinPatterns {
		copy(parr[:], pslice) // convert slice to array
		win, _ := IsWinnable(board, symbol, parr)
		if win {
			return win, parr
		}
	}

	return false, parr //can't win
}

func IsFree(board string, pos int) bool {

	switch {
	case pos > 10:
		return false
	case pos < 0:
		return false
	case pos == 3:
		return false
	case pos == 7:
		return false
	case board[pos] == '-':
		return true
	default:
		return false
	}
}

func MakeMove(board string, pos int, symbol byte) (string, error) {
	boardBytes := []byte(board)

	if board[pos] != '-' {
		return "", fmt.Errorf("pos %v is not free", pos)
	}

	boardBytes[pos] = symbol
	return string(boardBytes), nil
}

func GetMoveDifference(prev string, curr string) (int, error) {
	var diffCount int

	if !IsValidBoard(prev) {
		return 0, fmt.Errorf("invalid prev board %v", prev)
	} else if !IsValidBoard(curr) {
		return 0, fmt.Errorf("invalid curr board %v", curr)
	}

	for i := 0; i < len(curr); i++ {
		if prev[i] != curr[i] {
			diffCount++
		}
	}

	return diffCount, nil
}

func MakeWinMove(board string, move [3]int, symbol byte) (string, error) {
	boardBytes := []byte(board)

	winnable, _ := IsWinnable(board, symbol, move)
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
	winnable, _ := IsWinnable(board, getOpponentSymbol(symbol), move)
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

	emptyPositions := getEmptyPosList(board)
	if emptyPositions == nil {
		return "", fmt.Errorf("MakeRandomMove(): no empty positions")
	}

	rand.Seed(200)

	pos := emptyPositions[rand.Intn(len(emptyPositions))]
	boardBytes[pos] = symbol

	return string(boardBytes), nil
}

func getEmptyPosList(board string) []int {
	var emptyPos []int
	for i, ch := range board {
		if ch == '-' {
			emptyPos = append(emptyPos, i)
		}
	}
	return emptyPos
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

	for _, index := range WinPatterns {
		// fmt.Println(index)
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
