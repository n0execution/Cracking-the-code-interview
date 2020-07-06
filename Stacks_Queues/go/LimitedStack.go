package types

import (
  "fmt"
)


// Stack implementation
type LimitedStack struct {
  maxSize int
  size int
  elements []int
  top int
}


func (s *LimitedStack) Push(x int) bool {
  if s.IsFull() {
    return false
  }

  s.elements = append(s.elements, x)
  s.top = x
  s.size++
  return true
}


func (s *LimitedStack) Pop() (int, bool) {
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


func (s *LimitedStack) IsEmpty() bool {
  return s.size == 0
}


func (s *LimitedStack) IsFull() bool {
  return s.size == s.maxSize
}


func PrintLimitedStack(s LimitedStack) {
  fmt.Println(s.elements)
}


func main() {
  s := LimitedStack{maxSize:5}
  s.Push(1)
  s.Push(2)
  s.Push(3)
  s.Push(4)
  s.Push(5)
  s.Push(6)
  s.Push(6)

  PrintLimitedStack(s)
}
