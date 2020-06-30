package main

import (
  "types"
  "math/rand"
)


func delete(node *types.Node) {
  node.Data = node.Next.Data
  node.Next = node.Next.Next
}


func performPartition(list types.LinkedList, partition int) types.LinkedList {
  node := list.Head

  for node.Next != nil {
    if node.Data < partition && node != list.Head {
      list.AddToBeginning(node.Data)
      delete(node)
    } else {
      node = node.Next
    }
  }

  if node.Data < partition && node != list.Head {
    list.AddToBeginning(node.Data)
    list.DeleteLast()
  }

  return list
}


func main() {
  var list types.LinkedList
  var length int = 5

  for i := 0; i < length; i++ {
    list.Append(rand.Intn(9))
  }

  types.PrintLinkedList(list)
  list = performPartition(list, 6)
  types.PrintLinkedList(list)
}
