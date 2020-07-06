package main

import (
	"types"
)


func insertTheSmallest(s *types.Stack, tempStack *types.Stack, temp int) {
  var top, topTemp, nums int
  for {
    if tempStack.IsEmpty() {
      break
    }

    topTemp, _ = tempStack.Peek()
    if temp >= topTemp {
      break
    }

    top, _ = tempStack.Pop()
    s.Push(top)
    nums++
  }

  tempStack.Push(temp)

  for i := 0; i < nums; i++ {
    top, _ = s.Pop()
    tempStack.Push(top)
  }
}


func Sort(s *types.Stack) {
	tempStack := &types.Stack{}
	var top, topTemp int

	for !s.IsEmpty() {
		top, _ = s.Pop()

		if tempStack.IsEmpty() {
			tempStack.Push(top)
		}

		topTemp, _ = tempStack.Peek()
		if top >= topTemp {
			tempStack.Push(top)
		} else {
			insertTheSmallest(s, tempStack, top)
		}
	}

	for !tempStack.IsEmpty() {
		top, _ = tempStack.Pop()
		s.Push(top)
	}

}

func main() {
	var s types.Stack
	s.Push(1)
	s.Push(-23)
	s.Push(0)
	s.Push(3)
	s.Push(2)
	s.Push(10)
	Sort(&s)
	types.PrintStack(s)
}
