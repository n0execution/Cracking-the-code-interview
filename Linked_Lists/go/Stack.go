package types

import (
  "fmt"
)


// Stack implementation
type Stack struct {
  elements []int
  top int
  size int
}


func (s *Stack) Push(x int) {
  s.elements = append(s.elements, x)
  s.top = x
  s.size++
}


func (s *Stack) Pop() (int, bool) {
  top := s.top

  if s.IsEmpty() {
    return 0, false
  }

  s.elements = s.elements[:s.size-1]
  s.size--

  if s.IsEmpty() {
    return top, true
  }
  s.top = s.elements[s.size-1]

  return top, true
}


func (s *Stack) IsEmpty() bool {
  return s.size == 0
}


func PrintStack(s Stack) {
  fmt.Println(s.elements)
}
