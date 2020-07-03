package main

import (
  "fmt"
)

// StacksInArray implementation
type StacksInArray struct {
  stacksNum int
  stackSize int
  elements []int
  lengths []int
}


func CreateStack(stacksNum int, stackSize int) StacksInArray {
  s := StacksInArray{}
  s.stacksNum = stacksNum
  s.stackSize = stackSize

  for i := 0; i < stackSize * stacksNum; i++ {
    s.elements = append(s.elements, 0)
  }

  for i := 0; i < stacksNum; i++ {
    s.lengths = append(s.lengths, 0)
  }

  return s
}


func (s *StacksInArray) Push(stackNum, x int) bool{
  var startIndex, index int

  if stackNum < 0 || stackNum >= s.stackSize {
    return false
  }

  if s.lengths[stackNum] == s.stackSize {
    return false
  }

  startIndex = stackNum * s.stackSize
  index = startIndex + s.lengths[stackNum]
  s.elements[index] = x
  s.lengths[stackNum]++

  return true
}


func (s *StacksInArray) Pop(stackNum int) (int, bool){
  var startIndex, index, top int

  if stackNum < 0 || stackNum >= s.stackSize {
    return0,  false
  }

  if s.isEmpty(stackNum) {
    return 0, false
  }

  startIndex = stackNum * s.stackSize
  index = startIndex + s.lengths[stackNum] - 1
  top = s.elements[index]
  s.elements[index] = 0
  s.lengths[stackNum]--

  return top, true
}


func (s *StacksInArray) isEmpty(stackNum int) bool {
  return s.lengths[stackNum] == 0
}


func main() {
  s := CreateStack(3, 3)

  s.Push(0, 3)
  s.Push(0, 35)

  s.Push(1, 1)
  s.Push(1, 2)
  s.Push(1, 3)

  fmt.Println(s.Pop(0))
  fmt.Println(s.Pop(1))
}
