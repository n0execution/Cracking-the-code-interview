package main

import (
  "types"
  "fmt"
)


func detect(node1 *types.Node) *types.Node {
  node2 := node1.Next

  for node1 != nil {
    if node1 == node2 {
      return node1
    }

    node1 = node1.Next
    node2 = node2.Next.Next
  }

  return &types.Node{}
}


func main() {
  var list types.LinkedList
  var loopBegin *types.Node

  list.Append(1)
  list.Append(2)
  list.Append(3)
  list.Append(4)
  node := list.Head


  for node.Next != nil {
    node = node.Next
    loopBegin = node
  }

  list.Append(5)
  list.Append(6)
  list.Append(7)

  node = list.Head

  for node.Next != nil {
    node = node.Next
  }

  node.Next = loopBegin

  fmt.Println(detect(list.Head))
}
