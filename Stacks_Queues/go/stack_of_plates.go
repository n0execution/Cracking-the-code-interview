package main
import (
  "types"
  "fmt"
)


type SetOfStacks struct {
  MaxStackSize int
  stacks []*types.LimitedStack
  stacksNum int
}


func (set *SetOfStacks) Push(x int) {
  var stack *types.LimitedStack
  if set.stacksNum == 0 {
    stack = &types.LimitedStack{MaxSize: set.MaxStackSize}
    set.stacks = append(set.stacks, stack)
    set.stacksNum++
  } else {
    stack = set.stacks[set.stacksNum - 1]
  }

  if stack.IsFull() {
    stack := types.LimitedStack{MaxSize: set.MaxStackSize}
    stack.Push(x)
    set.stacks = append(set.stacks, &stack)
    set.stacksNum++
  } else {
    stack.Push(x)
  }
}


func (set *SetOfStacks) Pop() (int, bool) {
  return set.PopAt(set.stacksNum - 1)
}


func (set *SetOfStacks) PopAt(stackNum int) (int, bool) {
  var stack *types.LimitedStack
  if set.stacksNum == 0 {
    return 0, false
  }

  if stackNum < 0 || stackNum >= set.stacksNum {
    return 0, false
  }

  stack = set.stacks[stackNum]
  top, _ := stack.Pop()

  if stack.IsEmpty() {
    if stackNum == set.stacksNum - 1 {
      set.stacks = set.stacks[:set.stacksNum-1]
    } else {
      set.stacks = append(set.stacks[:stackNum], set.stacks[stackNum + 1:]...)
    }
    set.stacksNum--
  }

  return top, true
}

func (set *SetOfStacks) PrintSetOfStacks() {
  for i := 0; i < set.stacksNum; i++ {
    types.PrintLimitedStack(*set.stacks[i])
  }
  fmt.Println("\n")
}


func main() {
  s := SetOfStacks{MaxStackSize: 2}

  s.Push(3)
  s.Push(2)
  s.Push(1)
  s.Push(2)
  s.PrintSetOfStacks()
  s.Push(3)
  s.Push(4)
  s.PrintSetOfStacks()
  fmt.Println(s.PopAt(1))
  fmt.Println(s.PopAt(1))
  s.PrintSetOfStacks()
}
