package tictactoe

import (
	"strings"
	"testing"
	"reflect"
)

func TestGetEmptyBoard(t *testing.T) {
	const EXPECTED = "---|---|---"

	empBoard := GetEmptyBoard()
	if strings.Compare(empBoard, EXPECTED) != 0 {
		t.Error("GetEmptyBoard() returned %v", empBoard)
	}
}

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

func TestGetEmptyPos(t *testing.T)  {
	tests := []struct	{
		board string
		want []int
	}	{
		{"OX-|-OO|X-X", []int{2, 4, 9}},
		{"---|---|---", []int{0,1,2,4,5,6,8,9,10}},
		{"OXO|XOX|XXO", nil},
	}

	for _, test := range tests	{
		got := getEmptyPosList(test.board)
		if !reflect.DeepEqual(got, test.want)	{
			t.Errorf("getEmptyPosList(%v) got=%v want=%v", test.board, got, test.want)
		}
	}
}

func TestIsWinnable(t *testing.T) {
	tests := []struct {
		brd  string
		sym  byte
		ptrn [3]int
		want bool
	}{
		{"X-X|-OO|X-X", 'O', [3]int{1, 2, 3}, false},
		{"X-X|-OO|X-X", 'O', [3]int{4, 5, 6}, true},
		{"X-X|-OO|X-X", 'X', [3]int{0, 1, 2}, true},
		{"X-X|-OO|X-X", 'X', [3]int{8,9,10}, true},

		{"X-X|-OO|X-X", 'X', [3]int{1, 2, 3}, false},
		{"X-X|-OO|X-X", 'X', [3]int{0, 4, 8}, true},
		{"X-X|-OO|X-X", 'X', [3]int{2, 6, 10}, false},
		{"X-X|-OO|X-X", 'X', [3]int{5, 5, 5}, false},
		{"X-X|-OO|X-X", 'X', [3]int{3, 4, 5}, false},
		{"X-X|-OO|X-X", 'X', [3]int{0, 0, 0}, false},
	}

	for _, test := range tests {
		if got, _ := IsWinnable(test.brd, test.sym, test.ptrn); got != test.want {
			t.Errorf("IsWinnable(%q, %q, %v) want:%v, got:%v", test.brd, test.sym, test.ptrn, test.want, got)
		}
	}
}

func TestCanWinNext(t *testing.T)  {
	var emptyArray [3]int
	tests := []struct	{
		brd string
		sym byte
		ptrn [3]int
		want bool
	}	{
			{"X-X|-OO|X-X", 'O', [3]int{4,5,6}, true},
			{"X-X|-OO|X-X", 'X', [3]int{0,1,2}, true},
			{"X--|---|X--", 'X', [3]int{0,4,8}, true},
			{"X--|---|---", 'O', emptyArray, false},
			{"X--|-X-|---", 'X', [3]int{0,5,10}, true},
			{"---|-O-|O--", 'O', [3]int{2,5,8}, true},
	}

	for _, test := range tests	{
		win, p := CanWinNext(test.brd, test.sym);
		if win != test.want	{
			t.Errorf("CanWinNext(%v, %q) want:%v, got:%v", test.brd, test.sym, test.want, win)
		}

		if win	{
			if p != test.ptrn	{
				t.Errorf("CanWinNext(%v, %q) want:%v, got:%v", test.brd, test.sym, test.ptrn, p)
			}
		}
	}
}

func TestMakeWinMove(t *testing.T) {
	winBoard := "X--|X--|---"
	winIndices := [3]int{0, 4, 8}

	const EXPECTED = "X--|X--|X--"

	finalBoard, _ := MakeWinMove(winBoard, winIndices, 'X')
	if strings.Compare(finalBoard, EXPECTED) != 0 {
		t.Errorf("MakeWinMove(%v, %v %q) expected:%v, actual:%v", winBoard, winIndices, 'X', EXPECTED, finalBoard)
	}
}

func TestBlockWinMove(t *testing.T) {
	winBoard := "X--|X--|---"
	winIndices := [3]int{0, 4, 8}

	const EXPECTED = "X--|X--|O--"
	finalBoard, _ := BlockWinMove(winBoard, winIndices, 'O')
	if strings.Compare(finalBoard, EXPECTED) != 0 {
		t.Errorf("BlockWinMove(%v, %v %q) expected:%v, actual:%v", winBoard, winIndices, 'X', EXPECTED, finalBoard)
	}
}

func TestMakeRandomMove(t *testing.T) {
	someBoard := "X--|O--|-O-"
	var symbol byte = 'X'
	finalBoard, _ := MakeRandomMove(someBoard, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, symbol)
	if strings.Compare(someBoard, finalBoard) == 0 {
		// no move happened
		t.Errorf("MakeRandomMove(%v, %q) returned identical %v", someBoard, symbol, finalBoard)
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

func TestGetMoveDifference(t *testing.T) {
	var tests = []struct {
		prev string
		curr string
		want int
	}{
		{"---|---|---", "---|---|---", 0},
		{"---|---|---", "---|---|---", 0},
		{"---|---|---", "--X|---|---", 1},
		{"---|---|---", "--O|---|---", 1},
		{"---|---|---", "--X|---|--X", 2},
		{"XX-|---|---", "XX-|-O-|---", 1},
		{"XX-|---|---", "XO-|-OO|-O-", 4},
	}

	for _, test := range tests {
		if got, _ := GetMoveDifference(test.prev, test.curr); got != test.want {
			t.Errorf(" GetMoveDifference(%q,%q) = %v", test.prev, test.curr, got)
		}
	}
}
