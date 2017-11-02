package tictactoe

import "testing"

func TestIsValid(t *testing.T) {
	validBoards := [...]string{"X--|XXX|OXO",
		"X--|XXX|OXO",
		"XXX|XXX|XXX",
		"x--|oxo|oox",
		"---|---|---",
		"X--|XXX|OXo",
		"OOO|OOO|OOO",
	}

	for _, validBoard := range validBoards {
		if !IsValidBoard(validBoard) {
			t.Error(" returns invalid for valid board ", validBoard)
		}
	}

	invalidBoards := [...]string{"abc",
		"X--XXXOXO",
		"XXX-XXX-XXX",
		"O-|-XX|XXO",
		"||",
	}

	for _, invalidBoard := range invalidBoards {
		if IsValidBoard(invalidBoard) {
			t.Error(" returns valid for invalid board ", invalidBoard)
		}
	}
}

func TestIsWinnable_Win(t *testing.T) {
	winnable_board := "X-X|-OO|X-X"

	// stores indices; true indicates winnable, false otherwise
	indices := map[[3]int]bool{
		{1, 2, 3}:  false,
		{0, 4, 8}:  true,
		{2, 6, 10}: false,
		{5, 5, 5}:  false,
		{3, 4, 5}:  false,
		{0, 0, 0}:  false,
	}

	for idxTriplet, expected := range indices {
		actual, _, _ := IsWinnable(winnable_board, 'X', idxTriplet)
		if expected != actual {
			t.Errorf("IsWinnable(%v, %q	, %v) expected: %v, actual: %v", winnable_board, 'X', idxTriplet, expected, actual)
		}
	}
}

func TestIsWinnable_WinReturn(t *testing.T) {
	winnable_board := "X-X|-OO|X-X"
	winning_indices := [3]int{0, 4, 8}
	const EXPECTEDLENGTH = 3

	winstatus, indices, _ := IsWinnable(winnable_board, 'X', winning_indices)

	if winstatus != true {
		t.Errorf("IsWinnable(%v, %q	, %v) expected: %v, actual: %v", winnable_board, 'X', winning_indices, true, winstatus)
	}

	if len(indices) != EXPECTEDLENGTH {
		t.Errorf("IsWinnable(%v, %q, %v) returned slice with len != %d (==%v)", winnable_board, 'X', winning_indices, EXPECTEDLENGTH, len(indices))
	}

	if indices[0] != 0 || indices[1] != 4 || indices[2] != 8 {
		t.Errorf("IsWinnable(%v, %q, %v) expected: %v, actual: %v", winnable_board, 'X', winning_indices, winning_indices, indices)
	}
}

func TestHasWon(t *testing.T) {
	winningBoards := map[string]byte{
		"X--|X--|X--": 'X',
		"XOO|XOO|XOO": 'X',
		"-X-|-X-|-X-": 'X',
		"--X|--X|--X": 'X',
		"XXX|---|---": 'X',
		"---|XXX|---": 'X',
		"---|---|XXX": 'X',
		"X--|-X-|--X": 'X',
		"--X|-X-|X--": 'X',

		"O--|O--|O--": 'O',
		"-O-|-O-|-O-": 'O',
		"XOX|XOX|XOX": 'O',
		"--O|--O|--O": 'O',
		"OOO|---|---": 'O',
		"---|OOO|---": 'O',
		"---|---|OOO": 'O',
		"O--|-O-|--O": 'O',
		"--O|-O-|O--": 'O',

		"XXX|XXX|XXX": 'X',
		"OOO|OOO|OOO": 'O',
	}

	wrongSymbolBoards := map[string]byte{
		"X--|X--|X--": 'O',
		"-X-|-X-|-X-": 'O',
		"--X|--X|--X": 'O',
		"XXX|---|---": 'O',
		"---|XXX|---": 'O',
		"---|---|XXX": 'O',
		"X--|-X-|--X": 'O',
		"--X|-X-|X--": 'O',

		"O--|O--|O--": 'X',
		"-O-|-O-|-O-": 'X',
		"--O|--O|--O": 'X',
		"OOO|---|---": 'X',
		"---|OOO|---": 'X',
		"---|---|OOO": 'X',
		"O--|-O-|--O": 'X',
		"--O|-O-|O--": 'X',
	}

	nonWinningBoards := map[string]byte{
		"-X-|---|-X-": 'X',
		"-X-|-O-|-X-": 'X',
		"-X-|--X|-X-": 'X',
	}

	for board, symbol := range winningBoards {
		won := HasWon(board, symbol)
		if won != true {
			t.Errorf("HasWon(%v, %q) returns %v", board, symbol, won)
		}
	}

	for board, symbol := range wrongSymbolBoards {
		won := HasWon(board, symbol)
		if won == true {
			t.Errorf("HasWon(%v, %q) returns %v", board, symbol, won)
		}
	}

	for board, symbol := range nonWinningBoards {
		won := HasWon(board, symbol)
		if won == true {
			t.Errorf("HasWon(%v, %q) returns %v", board, symbol, won)
		}
	}
}

// func TestIsWinnable_All(t *testing.T)  {
//   row_indices := [][]int  {
//     {0,1,2},
//     {4,5,6},
//     {8,9,10},
//   }
//
//   col_indices := [][]int {
//     {0, 4, 8},
//     {1, 5, 9},
//     {2, 6, 10},
//   }
//
//   diag_indices := [][]int {
//     {0, 5, 10},
//     {2, 5, 8},
//   }
//
//
//
// }

func TestIsValidIndex(t *testing.T) {
	indexes := map[int]bool{
		-10: false,
		-1:  false,
		0:   true,
		1:   true,
		5:   true,
		9:   true,
		10:  true,
		11:  false,
		15:  false,
	}

	for key, expected := range indexes {
		if isValidIndex(key) != expected {
			t.Errorf(" isValidIndex[%d] expected:%v, actual:%v", key, expected, isValidIndex(key))
		}
	}
}

func TestIsIndicesDifferent(t *testing.T) {
	indices := map[[3]int]bool{
		{1, 2, 3}: true,
		{5, 5, 5}: false,
		{2, 2, 3}: false,
		{4, 5, 4}: false,
		{6, 4, 4}: false,
	}

	for idxTriplet, expected := range indices {
		actual := isIndicesDifferent(idxTriplet)

		if actual != expected {
			t.Errorf(" isIndicesDifferent[%v] expected:%v, actual:%v", idxTriplet, expected, actual)
		}
	}
}

func TestGetOpponentSymbol(t *testing.T) {

	symbolTable := map[byte]byte{
		'X': 'O',
		'O': 'X',
		'x': '-',
		'o': '-',
		'k': '-',
		'-': '-',
	}

	for sym, exp := range symbolTable {
		act := getOpponentSymbol(sym)

		if exp != act {
			t.Errorf("getOpponentSymbol(%q) expected:%q actual:%q", sym, exp, act)
		}

	}
}
