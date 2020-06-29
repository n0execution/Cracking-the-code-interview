package main

import (
  // "fmt"
  "types"
  "math/rand"
)


func valueInArray(array []int, value int) bool {
  if len(array) == 0 {
    return false
  }

  for _, num := range array {
    if num == value {
      return true
    }
  }
  return false
}


func removeDuplicates(list types.LinkedList) {
  nums := make([]int, 0)
  node := list.Head
  prev := node

  for node != nil {
    if valueInArray(nums, node.Data) {
      prev.Next = node.Next
    } else {
      nums = append(nums, node.Data)
      prev = node
    }
    node = node.Next
  }

}


func main() {
  var list types.LinkedList
  var length int = 5

  for i := 0; i < length; i++ {
    list.Append(rand.Intn(23))
  }

  list.Append(0)
  list.Append(1)
  list.Append(1)
  list.Append(22)
  list.Append(4)
  list.Append(5)
  list.Append(4)

  types.PrintLinkedList(list)
  removeDuplicates(list)
  types.PrintLinkedList(list)
}
