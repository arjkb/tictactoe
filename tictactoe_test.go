package tictactoe

import "testing"

func TestIsValid_valid(t *testing.T) {
	validBoards := [...]string{"X--|XXX|OXO",
		"X--|XXX|OXO",
		"XXX|XXX|XXX",
		"OOO|OOO|OOO",
		"---|---|---"}

	for _, validBoard := range validBoards {
		if !isValidBoard(validBoard) {
			t.Error(" Fails for valid board ", validBoard)
		}
	}
}
