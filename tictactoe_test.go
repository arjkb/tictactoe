package tictactoe

import "testing"

func TestFoo(t *testing.T)  {

  s := foo(100);
  if s != 10000 {
    t.Error("Whoa!")
  }

}
