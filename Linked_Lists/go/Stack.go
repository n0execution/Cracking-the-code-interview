package types

import (
  "fmt"
)


// Stack implementation
type Stack struct {
  elements []int
  top int
}


func (s *Stack) Push(x int) {
  s.elements = append(s.elements, x)
  s.top = x
}


func (s *Stack) Pop() (int, bool) {
  top := s.top
  length := len(s.elements)

  if length == 0 {
    return 0, false
  }

  s.elements = s.elements[:length-1]

  if length == 1 {
    return top, true
  }
  s.top = s.elements[length-2]

  return top, true
}


func PrintStack(s Stack) {
  fmt.Println(s.elements)
}
