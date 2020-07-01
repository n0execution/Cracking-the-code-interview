package main

import (
  "fmt"
  "types"
)


func compareLists(node1 *types.Node, node2 *types.Node) bool {
  for node1.Next != nil {
    if node1.Data != node2.Data {
      return false
    }

    node1 = node1.Next
    node2 = node2.Next
  }

  return true
}


func reverseLinkedList(node *types.Node) types.LinkedList {
  if node.Next == nil {
    var resultList types.LinkedList
    resultList.Append(node.Data)

    return resultList
  }

  resultList := reverseLinkedList(node.Next)
  resultList.Append(node.Data)

  return resultList
}


func isPalindrome(list types.LinkedList) bool {
  reversedList := reverseLinkedList(list.Head)

  return compareLists(list.Head, reversedList.Head)
}


func main() {
  var list types.LinkedList

  list.Append(1)
  list.Append(2)
  list.Append(3)
  list.Append(4)
  list.Append(3)
  list.Append(2)
  list.Append(1)

  fmt.Println(isPalindrome(list))
}
