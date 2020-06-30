package main

import (
  "types"
  "math/rand"
  "fmt"
)


func kToLast(head *types.Node, k int) int {
  if head.Next == nil{
    return 0
  }

  index := kToLast(head.Next, k) + 1

  if index == k {
    fmt.Println(k, head.Data)
  }
  return index
}


func kToLast2(head *types.Node, k int) {
  p1 := head
  p2 := head

  for i := 0; i < k; i++ {
    p2 = p2.Next
  }

  for p2.Next != nil {
    p1 = p1.Next
    p2 = p2.Next
  }

  fmt.Println(k, p1.Data)

}


func main() {
  var list types.LinkedList
  var length int = 5

  for i := 0; i < length; i++ {
    list.Append(rand.Intn(23))
  }

  types.PrintLinkedList(list)
  kToLast2(list.Head, 1)
}
