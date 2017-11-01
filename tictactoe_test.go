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

// func TestIsWinnable(t *testing.T)  {
//   winnable_board := "X--|X--|-X-",
//
//   indices := [[0,1,2], [0,4,8], [2,4,6]]
//
//
// }

func TestIsValidIndex(t *testing.T)  {
  indexes := map[int]bool  {
    -10:false,
    -1: false,
    0:true,
    1:true,
    5:true,
    9:true,
    10:true,
    11:false,
    15:false,
  }

  for key, expected := range indexes  {
    if isValidIndex(key) != expected  {
      t.Errorf(" isValidIndex[%d] expected:%v, actual:%v", key, expected, isValidIndex(key))
    }
  }

}
