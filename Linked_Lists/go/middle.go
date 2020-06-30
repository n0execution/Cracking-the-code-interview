package main

import (
  "types"
  "math/rand"
  // "fmt"
)


// Length of linked list
var Length int


func deleteMiddle(node *types.Node) {
  node.Data = node.Next.Data
  node.Next = node.Next.Next
}


func getMiddle(head *types.Node) int {
  if head.Next == nil {
    Length++
    return Length
  }

  Length++
  index := getMiddle(head.Next) - 1

  if index == Length / 2 {
    deleteMiddle(head.Next)
  }
  return index
}


func main() {
  var list types.LinkedList
  var length int = 5

  for i := 0; i < length; i++ {
    list.Append(rand.Intn(23))
  }

  types.PrintLinkedList(list)
  getMiddle(list.Head)
  types.PrintLinkedList(list)
}
