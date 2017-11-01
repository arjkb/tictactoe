package tictactoe

func foo(num int) int {
  return num * num;
}

func isValidBoard(board string) bool {

  if len(board) != 11  {
    return false
  }

  return true
}
