package main

import (
  "types"
  // "fmt"
  "math"
)


func generateList(n int, reverse bool) types.LinkedList {
  var resultList types.LinkedList

  for n / 10 != 0 || n % 10 != 0 {
    remainder := n % 10

    if reverse {
      resultList.Append(remainder)
    } else {
      resultList.AddToBeginning(remainder)
    }

    n = n / 10
  }

  return resultList
}


func sumListReverse(node *types.Node, tens float64) int {
  if node.Next == nil {
    return node.Data * int(math.Pow(10., tens))
  }
  return sumListReverse(node.Next, tens + 1) + node.Data * int(math.Pow(10., tens))
}


func sumListsReverse(node1 *types.Node, node2 *types.Node) types.LinkedList {
  return generateList(sumListReverse(node1, 0.) + sumListReverse(node2, 0.), true)
}


func sumList(node *types.Node) (int, float64) {
  if node.Next == nil {
    return node.Data, 1.
  }

  value, tens := sumList(node.Next)

  return node.Data * int(math.Pow(10., tens)) + value, tens + 1
}


func sumLists(node1 *types.Node, node2 *types.Node) types.LinkedList {
  firstValue, _ := sumList(node1)
  secondValue, _ := sumList(node2)
  return generateList(firstValue + secondValue, false)
}


func main() {
  var list1, list2 types.LinkedList
  list1.Append(7)
  list1.Append(1)
  list1.Append(6)
  list1.Append(9)

  list2.Append(5)
  list2.Append(9)
  list2.Append(2)

  types.PrintLinkedList(sumListsReverse(list1.Head, list2.Head))

  types.PrintLinkedList(sumLists(list1.Head, list2.Head))
}
