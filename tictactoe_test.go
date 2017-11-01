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
