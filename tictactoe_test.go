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

func TestIsWinnable(t *testing.T)  {
  winnable_board := "X-X|-OO|X-X"

  // stores indices; true indicates winnable, false otherwise
  indices := map[[3]int]bool  {
    {1,2,3}:false,
    {0,4,8}:true,
    {2,6,10}:false,
    {5,5,5}:false,
    {3,4,5}:false,
    {0,0,0}:false,
  }

  for idxTriplet, expected := range indices {
    actual, _ := IsWinnable(winnable_board, 'X', idxTriplet)
    if expected != actual {
      t.Errorf("IsWinnable(%v, %v, %v) expected: %v, actual: %v", winnable_board, 'X',idxTriplet, expected, actual)
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

func TestIsIndicesDifferent(t *testing.T)  {
  indices := map[[3]int]bool  {
    {1,2,3}:true,
    {5,5,5}:false,
    {2,2,3}:false,
    {4,5,4}:false,
    {6,4,4}:false,
  }

  for idxTriplet, expected := range indices {
    actual := isIndicesDifferent(idxTriplet)

    if actual != expected {
      t.Errorf(" isIndicesDifferent[%v] expected:%v, actual:%v", idxTriplet, expected, actual)
    }
  }
}
