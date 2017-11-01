package tictactoe

import "testing"

func TestFoo(t *testing.T)  {

  s := foo(100);
  if s != 10000 {
    t.Error("Whoa!")
  }

}

func TestIsValid(t *testing.T)  {
  validBoards := [...]string{"X--|XXX|OXO", "X--|XXX|OXO"}

  for _, validBoard := range validBoards {
    if !isValidBoard(validBoard) {
      t.Error(" Fails for valid board ", validBoard)
    }
  }
}
