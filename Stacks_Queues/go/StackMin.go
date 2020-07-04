package main

import (
  "fmt"
  "types"
)


// Stack implementation
type StackMin struct {
  elements []int
  mins types.DoublyLinkedList
  top int
}


func (s *StackMin) Push(x int) {
  s.elements = append(s.elements, x)
  s.top = x

  if s.mins.Head == nil {
    s.mins.Append(x)
    return
  }

  if x <= s.mins.Tail.Data {
    s.mins.Append(x)
  }
}


func (s *StackMin) Pop() (int, bool) {
  top := s.top
  length := len(s.elements)

  if length == 0 {
    return 0, false
  }

  s.elements = s.elements[:length-1]

  if top == s.mins.Tail.Data {
    s.mins.DeleteLast()
  }

  if length != 1 {
    s.top = s.elements[length-2]
  }

  return top, true
}


func PrintStack(s StackMin) {
  fmt.Println(s.elements)
}


func (s *StackMin) Min() int {
  return s.mins.Tail.Data
}


func main() {
  var s StackMin
  s.Push(34)
  s.Push(2)
  s.Push(32)
  s.Push(-3)
  s.Pop()
  s.Pop()

  PrintStack(s)
  fmt.Println(s.Min())
}
